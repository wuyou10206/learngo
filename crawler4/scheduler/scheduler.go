package scheduler

import "learngo/crawler4/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) GetWork() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request dowm to worker chan
	go func() {
		s.workerChan <- r
	}()
}
