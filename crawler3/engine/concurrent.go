package engine

import (
	"github.com/gpmgo/gopm/modules/log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}
type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	//in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkCount; i++ {
		createWorker(out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Warn("Got item #%d:%v", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}
func createWorker(out chan ParseResult, scheduler Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			// tell scheduler i'm ready
			scheduler.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
