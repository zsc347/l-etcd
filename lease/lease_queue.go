package lease

import "container/heap"

// LeaseWithTime contains lease object with a time.
// For the lessor's lease heap, time identifies the lease expiration time.
// For the lessor's lease checkpoint heap, the time identifies the
// next lease checkpoint time.
type LeaseWithTime struct {
	id LeaseID
	// Unix nanos timestamp
	time  int64
	index int
}

// LeaseQueue is priority queue for lease
type LeaseQueue []*LeaseWithTime

func (pq LeaseQueue) Len() int {
	return len(pq)
}

func (pq LeaseQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq LeaseQueue) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}

func (pq *LeaseQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*LeaseWithTime)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *LeaseQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// LeaseExpiredNotifier is a queue used to notify lessor
// to revoke expired lease.
// Only save one item for lease, `Register` will update time of the corresponding lease.
type LeaseExpiredNotifier struct {
	m     map[LeaseID]*LeaseWithTime
	queue LeaseQueue
}

func newLeaseExpiredNotifier() *LeaseExpiredNotifier {
	return &LeaseExpiredNotifier{
		m:     make(map[LeaseID]*LeaseWithTime),
		queue: make(LeaseQueue, 0),
	}
}

func (mq *LeaseExpiredNotifier) Init() {
	heap.Init(&mq.queue)
	mq.m = make(map[LeaseID]*LeaseWithTime)
	for _, item := range mq.queue {
		mq.m[item.id] = item
	}
}

func (mq *LeaseExpiredNotifier) RegisterOrUpdate(item *LeaseWithTime) {
	if old, ok := mq.m[item.id]; ok {
		old.time = item.time
		heap.Fix(&mq.queue, old.index)
	} else {
		heap.Push(&mq.queue, item)
		mq.m[item.id] = item
	}
}

func (mq *LeaseExpiredNotifier) Unregister() *LeaseWithTime {
	item := heap.Pop(&mq.queue).(*LeaseWithTime)
	delete(mq.m, item.id)
	return item
}

func (mq *LeaseExpiredNotifier) Poll() *LeaseWithTime {
	if mq.Len() == 0 {
		return nil
	}
	return mq.queue[0]
}

func (mq *LeaseExpiredNotifier) Len() int {
	return len(mq.m)
}
