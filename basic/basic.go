package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
	dd bool
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
	fmt.Printf("%d,%q\n", a, s)
}
func variableInitiaValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func variableTypeDeduction() {
	var a, b, c, d = 12, 1.3, true, "qwe"
	var s string
	fmt.Println(a, b, c, d, s)
}
func variableShorter() {
	a, b, c, d := 1, 4.5, false, "zxc"
	s := "aaa"
	fmt.Println(a, b, c, d, s)
}
func euler() {
	c := 3 + 4i
	x := cmplx.Abs(c)
	fmt.Println(x)
	n := cmplx.Pow(math.E, 1i*math.Pi)
	fmt.Println(n + 1)
	n2 := cmplx.Exp(1i * math.Pi)
	fmt.Println(n2 + 1)
	fmt.Printf("%.3f\n", n+1)
}
func triangle() {
	var a, b int = 3, 4
	var c int
	c = (int)(math.Sqrt((float64)(a*a + b*b)))
	fmt.Println(c)
}
func calTriangle(a, b int) int {
	var c int
	c = (int)(math.Sqrt((float64)(a*a + b*b)))
	return c
}
func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = (int)(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
	const (
		d = 5
		e = 6
	)
}
func enums() {
	//const (
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, javascript, python, golang)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitiaValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb, dd)

	euler()
	triangle()
	consts()
	enums()
}
