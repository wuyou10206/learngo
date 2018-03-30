package scheduler

import "learngo/crawler2/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigWorkerChan(in chan engine.Request) {
	s.workerChan = in
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request dowm to worker chan
	go func() {
		s.workerChan <- r
	}()
}
