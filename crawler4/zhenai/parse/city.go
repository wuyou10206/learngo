package parse

import (
	"learngo/crawler4/engine"
	"regexp"
)

var cityReg = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]+)</a></th>`)
var cityUrlReg = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/yulin/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {

	match := cityReg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		name := string(m[2])
		//	result.Items = append(result.Items, "User "+name)
		request := engine.Request{
			Url: string(m[1]),
			ParseFunc: func(contents []byte) engine.ParseResult {
				return ParsePerson(contents, name)
			},
		}
		result.Requests = append(result.Requests, request)
	}
	match = cityUrlReg.FindAllSubmatch(contents, -1)
	for _, m := range match {
		request := engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}
