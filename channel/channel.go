package main

import (
	"fmt"
	"time"
)

func main() {
	closeChannel()
}
func channelDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
func worker(id int, c chan int) {
	for {
		n := <-c
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
func worker2(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("worker %d received %d\n", id, n)

	}
}
func worker3(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func createWorker2(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Printf("worker %d received %d\n", id, n)
		}
	}()
	return c
}
func channelDemo2() {
	c := make(chan int)
	go worker(0, c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
func channelDemo3() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for i, c := range channels {
		c <- 'a' + i
	}
	for i, c := range channels {
		c <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
func channelDemo4() {
	c := createWorker(0)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
func channelDemo5() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i, c := range channels {
		c <- 'a' + i
	}
	for i, c := range channels {
		c <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
func channelDemo6() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker2(i)
	}
	for i, c := range channels {
		c <- 'a' + i
	}
	for i, c := range channels {
		c <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	go worker(0, c)
	time.Sleep(time.Millisecond)
}
func bufferedChannel2() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	time.Sleep(time.Millisecond)
}
func closeChannel() {
	c := make(chan int, 3)
	go worker3(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
	time.Sleep(time.Millisecond)
}
