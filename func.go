package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a,b int,op string) int{
	switch op {
	case "+":
		return a+b
	case "-":
		return a-b
	case "*":
		return a*b
	case "/":
		q,_:=div(a,b)
		return q
	default:
		panic("no support op:"+op)
	}
}
func eval2(a,b int,op string) (int,error){
	switch op {
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		q,_:=div(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf("no support op:%s",op)
	}
}
func apply(op func(int,int) int,a,b int) int{
	p:=reflect.ValueOf(op).Pointer()
	opName:=runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d)\n",opName,a,b)
	return op(a,b)
}
func sum(numbers ...int) int {
	s:=0
	for i:= range numbers{
		s+=numbers[i]
	}
	return s
}
func div(a,b int)(int,int){
	x:=a/b
	y:=a%b
	return x,y
}
func div2(a,b int)(q,r int){
	q=a/b
	r = a%b
	return
}
func swap(a,b *int){
	*b,*a = *a,*b
}
func swap2(a,b int)(int,int){
	b,a = a,b
	return b,a
}
func pow(a,b int) int{
	return (int)(math.Pow((float64)(a),(float64)(b)))
}
func main() {
	fmt.Println(eval(3,4,"*"))
	fmt.Println(div(13,3))
	fmt.Println(div2(13,3))
	q,r := div2(14,3)
	fmt.Println(q,r)
	fmt.Println(eval2(3,4,"x"))
	result,error:=eval2(3,4,"x");
	if error!=nil{
		fmt.Println(error)
	}else{
		fmt.Println(result)
	}
	fmt.Println(apply(pow,3,4))
	fmt.Println(apply(func(a int, b int) int {
		return a+b
	},3,4))
	fmt.Println(sum(1,2,3,4,5))
	a,b:=3,4
	swap(&a,&b)
	fmt.Println(a,b)
	fmt.Println(swap2(a,b))
}
