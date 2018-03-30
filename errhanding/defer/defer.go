package main

import (
	"bufio"
	"errors"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func trydefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}
func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
func writeFile2(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is error")
	if err != nil {
		fmt.Println("file already exists")
		fmt.Println("Error:" + err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
		//panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			break
		}
	}
}
func main() {
	//trydefer()
	tryDefer2()
	writeFile("fib.txt")
	writeFile2("fib.txt")
}
