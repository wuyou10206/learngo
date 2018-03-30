package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello world %d\n", i)
			}
		}(i)
	}
	//var a [10]int
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		for {
	//			a[i]++
	//			runtime.Gosched()
	//		}
	//	}(i)
	//
	//}
	time.Sleep(time.Minute)
	//	fmt.Println(a)
}
