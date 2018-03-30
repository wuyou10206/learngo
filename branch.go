package main

import (
	"io/ioutil"
	"fmt"
)
func grade(score int) string{
	g:=""
	switch  {
	case score<0||score>100:
		panic(fmt.Sprintf("Wrong score:%d",score))
	case score<60:
		g = "F"
	case score<80:
		g = "C"
	case score<90:
		g = "B"
	case score<=100:
		g = "A"
	}
	return g;
}
func main() {
	const filename = "abc.txt"
	contents,error := ioutil.ReadFile(filename)
	if error!=nil {
		fmt.Println(error)
	}else{
		fmt.Printf("%s\n",contents)
	//	fmt.Println(contents)
	}
	if contents2,error2:=ioutil.ReadFile(filename);error2!=nil{
		fmt.Println(error2)
	}else{
		fmt.Printf("%s\n",contents2)
	}
	grade(101)
	grade(0)
	grade(-1)
	grade(59)
	grade(88)
}
