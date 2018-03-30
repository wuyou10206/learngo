package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1500)))
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func main() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)
	var values []int
	n := 0
	//	hasValue := false
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var aW chan int
		var aV int
		if len(values) > 0 {
			aW = w
			aV = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		//	hasValue = true
		//w <- n
		case n = <-c2:
			values = append(values, n)
		//	hasValue = true
		//w <- n
		case aW <- aV:
			values = values[1:]
		case <-tick:
			fmt.Println(len(values))
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out")
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
