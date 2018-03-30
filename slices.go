package main

import "fmt"

func main() {
	arr:=[...]int{0,1,2,3,4,5,6,7}
	s:= arr[2:6]
	fmt.Println(s)
	fmt.Println(arr[:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:])
//	fmt.Println(arr[])
	s1:=arr[2:]
	s2:=arr[:]
	fmt.Println(s1)
	fmt.Println(s2)
	updateSlices(s1)
	fmt.Println("====")
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println(s2)
	updateSlices(s2)
	fmt.Println("----")
	fmt.Println(s1)
	fmt.Println(arr)
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
	arr = [...]int{0,1,2,3,4,5,6,7}
	s1=arr[2:6]
	s2 = s1[3:5]
	fmt.Println(s1,s2)
//	fmt.Println(s1[4])
//	fmt.Println(s1[3:10])
	fmt.Println(s1[:])
	fmt.Println(len(s1),cap(s1))
	fmt.Println(len(s2),cap(s2))
	fmt.Println(s2)
	s3:=append(s2,10)
	s4:=append(s3,11)
	s5:=append(s4,12)
	fmt.Println(s2,s3,s4,s5)
	fmt.Println(arr)
}
func updateSlices(s []int){
	s[0] = 100
}

