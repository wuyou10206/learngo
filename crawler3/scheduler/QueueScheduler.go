package scheduler

import "learngo/crawler3/engine"

type QueueScheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.workChan <- w
}

func (q *QueueScheduler) ConfigWorkerChan(in chan engine.Request) {
	panic("implement me")
}

func (q *QueueScheduler) Run() {
	q.workChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
