package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}
func main() {
	a := adder()
	f := a
	g := adder2(0)
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
		var v int
		v, g = g(i)
		fmt.Println(v)
	}
}
