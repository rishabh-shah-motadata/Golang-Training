package day4

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type pool struct {
	jobQueue chan job
	wg       sync.WaitGroup
	done     atomic.Bool
}

type job struct {
	data int
}

func newPool(workers int, jobs int) *pool {
	pool := &pool{
		jobQueue: make(chan job, jobs),
		wg:       sync.WaitGroup{},
	}
	pool.startWorkers(workers)

	pool.wg.Go(pool.autoScalePool)
	return pool
}

func (p *pool) startWorkers(workers int) {
	for range workers {
		p.wg.Go(func() {
			for job := range p.jobQueue {
				time.Sleep(300 * time.Millisecond)
				fmt.Println("processed data:", job.data*job.data)
			}
		})
	}
}

func (p *pool) autoScalePool() {
	for {
		switch {
		case p.done.Load():
			return
		case len(p.jobQueue) < cap(p.jobQueue)*3/4:
			p.startWorkers(10)
			time.Sleep(500 * time.Millisecond)
		case len(p.jobQueue) > cap(p.jobQueue)/2:
			p.startWorkers(5)
			time.Sleep(500 * time.Millisecond)
		default:
		}
	}
}

func Day4() {
	pool := newPool(10, 100)

	for i := range 200 {
		pool.jobQueue <- job{data: i}
	}
	close(pool.jobQueue)
	pool.done.Store(true)

	pool.wg.Wait()
}
