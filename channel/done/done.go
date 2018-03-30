package main

import (
	"fmt"
	"sync"
)

func main() {
	channelDemo4()
}
func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		go func() {
			w.done <- true
		}()
	}

}
func doWorker2(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		//	go func() {
		w.done <- true
		//	}()
	}

}
func doWorker3(id int, w worker2) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		//	go func() {
		w.wg.Done()
		//	}()
	}

}

func doWorker4(id int, w worker3) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		//	go func() {
		w.done()
		//	}()
	}

}

type worker2 struct {
	in chan int
	wg *sync.WaitGroup
}
type worker3 struct {
	in   chan int
	done func()
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker3(id int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		in: make(chan int),
		wg: wg,
	}
	go doWorker3(id, w)
	return w
}

func createWorker4(id int, wg *sync.WaitGroup) worker3 {
	w := worker3{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker4(id, w)
	return w
}

func createWorker2(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker2(id, w)
	return w
}
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w)
	return w
}
func channelDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	for i, w := range workers {
		w.in <- 'a' + i
		//<-w.done
	}
	for i, w := range workers {
		w.in <- 'A' + i
		//<-w.done
	}
	for _, w := range workers {
		<-w.done
		<-w.done
	}
}
func channelDemo4() {
	var workers [10]worker3
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createWorker4(i, &wg)
	}

	for i, w := range workers {
		w.in <- 'a' + i
		//<-w.done
	}
	for i, w := range workers {
		w.in <- 'A' + i
		//<-w.done
	}
	wg.Wait()
}
func channelDemo3() {
	var workers [10]worker2
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i] = createWorker3(i, &wg)
	}

	for i, w := range workers {
		w.in <- 'a' + i
		//<-w.done
	}
	for i, w := range workers {
		w.in <- 'A' + i
		//<-w.done
	}
	wg.Wait()
}
func channelDemo2() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i)
	}
	for i, w := range workers {
		w.in <- 'a' + i
		//<-w.done
	}
	for _, w := range workers {
		<-w.done
	}
	for i, w := range workers {
		w.in <- 'A' + i
		//<-w.done
	}
	for _, w := range workers {
		<-w.done
	}
}
