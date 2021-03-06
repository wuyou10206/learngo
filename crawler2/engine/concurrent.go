package engine

import "github.com/gpmgo/gopm/modules/log"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}
type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigWorkerChan(in)
	for i := 0; i < e.WorkCount; i++ {
		createWorker(in, out)
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
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
