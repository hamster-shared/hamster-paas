package service

import (
	"github.com/gammazero/workerpool"
)

// PoolService async goroutine util
type LongLinkPoolService struct {
	wp *workerpool.WorkerPool
}

func NewLongLinkPoolService() *LongLinkPoolService {
	wp := workerpool.New(100)

	return &LongLinkPoolService{
		wp: wp,
	}
}

func (l *LongLinkPoolService) Close() {
	l.wp.StopWait()
}

func (l *LongLinkPoolService) Submit(fn func()) {
	l.wp.Submit(fn)
}
