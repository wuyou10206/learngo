package main

import "fmt"

func main() {
	m:=map[string]string{
		"name":"ccmouse",
		"course":"golang",
		"site":"imooc",
		"quality":"notbad",
	}
	m2:=make(map[string]int)  //m2=empty map
	var m3 map[string]int   //m3=nil
	fmt.Println(m,m2,m3)
	for k,v:=range m{
		fmt.Println(k,v)
	}
	courseName,ok:=m["course"]
	fmt.Println(courseName,ok)
	if cName,ok:=m["cName"];ok{
		fmt.Println(cName,ok)
	}else{
		fmt.Println("key does not exist")
	}
	delete(m,"name")
	fmt.Println(m)

}
