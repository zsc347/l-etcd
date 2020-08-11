package lease

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/l-etcd/etcdserver/cindex"
	pb "github.com/l-etcd/etcdserver/etcdserverpb"
	"github.com/l-etcd/mvcc/backend"
	"go.uber.org/zap"
)

// LeaseID is type for lease
type LeaseID int64

// NoLease is a special LeaseID representing the absence of a lease.
const NoLease = LeaseID(0)

// MaxLeaseTTL is the maximum lease TTL value
const MaxLeaseTTL = 9000000000

var (
	forever = time.Time{}

	leaseBucketNmae = []byte("lease")

	// maximum number of leases to revoke per second; configurable for tests
	leaseRevokeRate = 1000

	// maximum number of lease checkpoints recorded to the consensus log per
	// second； configurable for tests
	leaseCheckpointRage = 1000

	// the default interval of lease checkpoint
	defaultLeaseCheckpointInterval = 5 * time.Minute

	// maximum number of lease checkpoints to batch into a single consensus log entry
	maxLeaseCheckpointBatchSize = 1000

	// the default interval to check if the expired lease is revoked
	defaultExpiredleaseRetryInterval = 3 * time.Second
)

// Errors of lease error
var (
	ErrNotPrimary       = errors.New("not a primary lessor")
	ErrLeaseNotFound    = errors.New("lease not found")
	ErrLeaseExists      = errors.New("lease already exists")
	ErrLeaseTTLTooLarge = errors.New("too large lease TTL")
)

// TxnDelete is a TxnWrite that only permits deltes. Defined here
// to avoid circular dependency with mvcc.
type TxnDelete interface {
	DeleteRange(key, end []byte) (n, rev int64)
	End()
}

// RangeDeleter is a TxnDelete constructor.
type RangeDeleter func() TxnDelete

// Checkpointer permits checkpointing of lease remaining TTLs to the consensus log.
// Defined here to avoid circular dependency with mvcc.
type Checkpointer func(ctx context.Context, lc *pb.LeaseCheckpointRequest)

// Lessor owns leases. It can grant, revoke, renew and modify leases for lessee.
type Lessor interface {
	// SetRangeDeleter lets the lessor create TxnDeletes to the store.
	// Lessor deletes the items in the revoked or expired lease by creating
	// new TxnDeletes.
	SetRangeDeleter(rd RangeDeleter)

	SetCheckPointer(cp Checkpointer)

	// Grant grants a lease that expires at least after TTL seconds.
	Grant(id LeaseID, ttl int64) (*Lease, error)

	// Revoke revokes a lease with given ID. The item attached to the
	// given lease will be removed. If the ID does not exist, an error
	// will be returned.
	Revoke(id LeaseID) error

	// Checkpoint applies the remainingTTL of a lease.
	// The remaining TTL is used in promote to set
	// the expiry of leases to less than the full TTL when possible.
	Checkpoint(id LeaseID, remainingTTL int64) error

	// Attach attaches given leaseItem to the lease with given LeaseID.
	// If the lease does not exist, an error will be returned.
	Attach(id LeaseID, items []LeaseItem) error

	// GetLease returns LeaseID for given item.
	// If no lease found, NoLease value will be returned.
	GetLease(item LeaseItem) LeaseID

	// Detach detaches given leaseItem from the lease with given LeaseID.
	// If the lease does not exist, an error will be returned.
	Detach(id LeaseID, items []LeaseItem) error

	// Promote promotes the lessor to be the primary lessor.
	// Primary lessor manages the expiration and renew of leases.
	// Newly promoted lessor renew the TTL of all lease to extend + previous TTL.
	Promote(extend time.Duration)

	// Demote demotes the lessor from being the primary lessor.
	Demote()

	// Renew renews a lease with given ID. It returns the renewed TTL.
	// If the ID does not exist, an error will be returned.
	Renew(id LeaseID) (int64, error)

	// Lookup gives the lease at a given lease id, if any
	Lookup(id LeaseID) *Lease

	// Leases lists all leases.
	Leases() []*Lease

	// ExpiredLeasesC returns a chan that is used to receive expired leases.
	ExpiredLeasesC() <-chan []*Lease

	// Recover recovers the lessor state from the given backend and RangeDeleter.
	Recover(b backend.Backend, rd RangeDeleter)

	// Stop stops the lessor for managing leases. The behavior of calling Stop multiple
	// times is undefined.
	Stop()
}

type lessor struct {
	mu sync.RWMutex

	// demotec is set when the lessor is the primary.
	// demotec will be closed if the lessor is demoted.
	demotec chan struct{}

	leaseMap             map[LeaseID]*Lease
	leaseExpiredNotifier *LeaseExpiredNotifier
	leaseCheckpopintHeap LeaseQueue
	itemMap              map[LeaseItem]LeaseID

	// When a lease expires, the lessor will delete the
	// leased range (or key) by the RangeDeleter.
	rd RangeDeleter

	// When a lease's deadline should be persisted to preserce the remaining TTL
	// across leader elections and restarts, the lessor will checkpoint the lease by
	// the Checkpointer.
	cp Checkpointer

	// backend to persist leases. We only persist lease ID and expiry for now.
	// The leased items can be recovered by iterating all the keys in kv.
	b backend.Backend

	// minLeaseTTL is the minimum lease TTL that can be granted for a lease.
	// Any requests for shorter TTLs are extended to the minimum TTL.
	minLeaseTTL int64

	expiredC chan []*Lease

	// stopC is a channel whose closure indicates that the lessor should be stopped.
	stopC chan struct{}

	// doneC is a channel whose closure indicates that the lessor is stopped.
	doneC chan struct{}

	lg *zap.Logger

	// Wait duration between lease checkpoints.
	checkpointInterval time.Duration

	// the interval to check if the expired lease is revoked
	expiredLeaseRetryInterval time.Duration

	ci cindex.ConsistentIndexer
}

type LessorConfig struct {
	MinLeaseTTL                int64
	CheckpointInterval         time.Duration
	ExpiredLeasesRetryInterval time.Duration
}

// func NewLessor(lg *zap.Logger,
// 	b backend.Backend,
// 	cfg LessorConfig,
// 	ci cindex.ConsistentIndexer) Lessor {
// 	return newLessor(lg, b, cfg, ci)
// }

func newLessor(lg *zap.Logger,
	b backend.Backend,
	cfg LessorConfig,
	ci cindex.ConsistentIndexer) *lessor {

	checkpointInterval := cfg.CheckpointInterval
	expiredLeaseRetryInterval := cfg.ExpiredLeasesRetryInterval
	if checkpointInterval == 0 {
		checkpointInterval = defaultLeaseCheckpointInterval
	}
	if expiredLeaseRetryInterval == 0 {
		expiredLeaseRetryInterval = defaultExpiredleaseRetryInterval
	}

	l := &lessor{
		leaseMap:                  make(map[LeaseID]*Lease),
		itemMap:                   make(map[LeaseItem]LeaseID),
		leaseExpiredNotifier:      newLeaseExpiredNotifier(),
		leaseCheckpopintHeap:      make(LeaseQueue, 0),
		b:                         b,
		minLeaseTTL:               cfg.MinLeaseTTL,
		checkpointInterval:        checkpointInterval,
		expiredLeaseRetryInterval: expiredLeaseRetryInterval,
		// expiredC is a small buffered chan to avoid unnecessary blocking.
		expiredC: make(chan []*Lease, 16),
		stopC:    make(chan struct{}),
		doneC:    make(chan struct{}),
		lg:       lg,
		ci:       ci,
	}

	l.initAndRecover()
	go l.runLoop()

	return l
}

func (le *lessor) initAndRecover() {
	// TODO
}

func (le *lessor) runLoop() {
	defer close(le.doneC)

	for {
		le.revokeExpiredLeases()
		le.checkpointScheduledLeases()

		select {
		case <-time.After(500 * time.Millisecond):
		case <-le.stopC:
			return
		}
	}
}

// revokeExpiredLeases finds all leases past their expiry and sends them
// to expired channel for to be revoked.
func (le *lessor) revokeExpiredLeases() {
	// TODO
}

// checkpointScheduledLeases finds all scheduled lease checkpoints that are due
// and submits them to the checkpointer to persist them to the consensus log.
func (le *lessor) checkpointScheduledLeases() {
	// TODO
}

type Lease struct {
	ID LeaseID
	// time to live of the lease in seconds
	ttl int64

	// remaining time to live in seconds,
	// if zero valued it is considered unset and the full ttl should be used
	remainingTTL int64

	// expiryMu protects concurrent accesses to expiry
	expiryMu sync.RWMutex

	expiry time.Time

	// mu protects concurrent accesses to itemSet
	mu      sync.RWMutex
	itemSet map[LeaseItem]struct{}
	revokec chan struct{}
}

type LeaseItem struct {
	Key string
}
