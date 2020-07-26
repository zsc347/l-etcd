package backend

import (
	"sync"
	"sync/atomic"
	"time"

	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
)

type backend struct {
	// size and commits are used with atomic operations to they must be
	// 64-bit aligned, otherwise 32-bit tests will crash

	// size is the number of bytes allocated in the backend
	size int64

	// sizeInUse is the number of bytes actually used in the backend
	sizeInUse int64

	// commits counts number of commits since start
	commits int64

	// openReadTxN is the number of currently open read transactions in the backend
	openReadTxN int64

	mu sync.RWMutex
	db *bolt.DB

	batchIntervel time.Duration
	batchLimit    int
	batchTx       *batchTxBuffered

	readTx *readTx

	stopc chan struct{}
	donec chan struct{}

	lg *zap.Logger
}

func (b *backend) begin(write bool) *bolt.Tx {
	b.mu.Lock()
	tx := b.unsafeBegin(write)
	b.mu.Unlock()

	size := tx.Size()
	db := tx.DB()
	stats := db.Stats()

	atomic.StoreInt64(&b.size, size)
	atomic.StoreInt64(&b.sizeInUse, size-(int64(stats.FreePageN))*int64(db.Info().PageSize))
	atomic.StoreInt64(&b.openReadTxN, int64(stats.OpenTxN))

	return tx
}

func (b *backend) unsafeBegin(write bool) *bolt.Tx {
	tx, err := b.db.Begin(write)
	if err != nil {
		b.lg.Fatal("failed to begin tx", zap.Error(err))
	}
	return tx
}
