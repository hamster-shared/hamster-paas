package service

import (
	"github.com/gammazero/workerpool"
	"runtime"
)

// PoolService async goroutine util
type PoolService struct {
	wp *workerpool.WorkerPool
}

func NewPoolService() *PoolService {
	numCPUs := runtime.NumCPU()

	wp := workerpool.New(numCPUs)

	return &PoolService{
		wp: wp,
	}
}

func (p *PoolService) Close() {
	p.wp.StopWait()
}

func (p *PoolService) Submit(fn func()) {
	p.wp.Submit(fn)
}
