package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2:=[3]int{2,3,4}
	arr3:=[...]int{1,2,3,4,5,6,7,8}
	var grid [4][5] bool
	arr1 = [5]int{1,2,3,4}
	arr1[4] = 12
	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)
	for i:=0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}
	for i:=range arr1{
		fmt.Println(arr1[i])
	}
	for i,v :=range arr3{
		fmt.Println(i,v)
	}
	for _,v:= range arr3{
		fmt.Println(v)
	}
	printArray(arr1)
	fmt.Println(arr1)
	printArray2(&arr1)
	fmt.Println(arr1)
	//printArray(arr2)
	printArray3(arr1[:])
	fmt.Println(arr1)
	arr4:=[...]int{1,2,3,4}
	arr5:=append(arr4[:],11,12)
	fmt.Println(arr5)

}
func printArray(arr [5]int){
	arr[0] = 100
	for i,v :=range arr{
		fmt.Println(i,v)
	}

}
func printArray2(arr *[5]int){
	arr[0] = 100
	for i,v :=range arr{
		fmt.Println(i,v)
	}

}
func printArray3(arr []int){
	arr[0] = 100
	for i,v :=range arr{
		fmt.Println(i,v)
	}

}