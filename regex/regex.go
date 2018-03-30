package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com@abc.com
email is abc@def.com
email2 is     kkk@qq.com
email3 is didi@did.com.cn
`

func main() {
	re := regexp.MustCompile(`(\w+)@([\w]+)(\.[\w\.]+)`)

	match := re.FindAllStringSubmatch(text, -1)
	//re.FindAllString(text, -1)
	//re.FindString(text)

	fmt.Println(match)

}
