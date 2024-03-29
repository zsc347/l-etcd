package schedule

import (
	"context"
	"sync"
)

// Job run by scheduler
type Job func(context.Context)

// Scheduler can schedule jobs
type Scheduler interface {
	// Schedule ask the scheduler to schedule a job defined by a given func.
	// Schedule to a stopped scheduler might panic.
	Schedule(j Job)

	// Pending return the number of pending jobs
	Pending() int

	// Scheduled returns the number of scheduled jobs (excluding pending jobs)
	Scheduled() int

	// Finished return the number of finished jobs
	Finished() int

	// WaitFinish wait util n jobs are finished and all pending jobs are finished
	WaitFinish(n int)

	// Stop stops the scheduler
	Stop()
}

type fifo struct {
	mu sync.Mutex

	resume     chan struct{}
	finishCond *sync.Cond
	donec      chan struct{}

	scheduled int
	finished  int

	pendings []Job

	ctx    context.Context
	cancel context.CancelFunc
}

// NewFIFOScheduler returns a Scheduler that schedules jobs in FIFO
// order sequentially
func NewFIFOScheduler() Scheduler {
	f := &fifo{
		resume: make(chan struct{}),
		donec:  make(chan struct{}),
	}
	f.finishCond = sync.NewCond(&f.mu)
	f.ctx, f.cancel = context.WithCancel(context.Background())

	go f.run()
	return f
}

// Schedule schedule a job that will be run in FIFO order sequentially
func (f *fifo) Schedule(j Job) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.cancel == nil {
		panic("schedule: schedule to stopped scheduler")
	}

	if len(f.pendings) == 0 {
		select {
		case f.resume <- struct{}{}:
		default:
		}
	}
	f.pendings = append(f.pendings, j)
}

func (f *fifo) Pending() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.pendings)
}

func (f *fifo) Scheduled() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.scheduled
}

func (f *fifo) Finished() int {
	f.finishCond.L.Lock()
	defer f.finishCond.L.Unlock()
	return f.finished
}

func (f *fifo) WaitFinish(n int) {
	f.finishCond.L.Lock()
	for f.finished < n || len(f.pendings) != 0 {
		f.finishCond.Wait()
	}
	f.finishCond.L.Unlock()
}

// Stop stops the scheduler and cancels all pending jobs.
func (f *fifo) Stop() {
	f.mu.Lock()
	f.cancel()
	f.cancel = nil
	f.mu.Unlock()
	<-f.donec
}

func (f *fifo) run() {
	defer func() {
		close(f.donec)
		close(f.resume)
	}()

	for {
		var todo Job

		f.mu.Lock()
		if len(f.pendings) != 0 {
			f.scheduled++
			todo = f.pendings[0]
		}
		f.mu.Unlock()

		if todo != nil {
			todo(f.ctx)
			f.finishCond.L.Lock()
			f.finished++
			f.pendings = f.pendings[1:]
			f.finishCond.Broadcast()
			f.finishCond.L.Unlock()
		} else {
			select {
			case <-f.resume:
			case <-f.ctx.Done():
				f.mu.Lock()
				pendings := f.pendings
				f.pendings = nil
				f.mu.Unlock()
				// clean up pending jobs
				for _, todo := range pendings {
					todo(f.ctx)
				}
				return
			}
		}
	}
}
