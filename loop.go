package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printFile(filename string) {
	file, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	printFileContents(file)
	//scanner:=bufio.NewScanner(file)
	//for scanner.Scan(){
	//	fmt.Println(scanner.Text())
	//}
}
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	fmt.Println(
		converToBin(13),
		converToBin(0),
		converToBin(123),
	)
	printFile("abc.txt")
	s := `abf
	asdd
	
	2323
	p
	`
	printFileContents(strings.NewReader(s))
	//	forever()
}
func forever() {
	for {
		fmt.Println("abc.txt")
	}
}
func converToBin(n int) string {
	str := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		str = strconv.Itoa(lsb) + str
	}
	return str
}
