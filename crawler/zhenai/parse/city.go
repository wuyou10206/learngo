package parse

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityReg = `<th><a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]+)</a></th>`

func ParseCity(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityReg)
	match := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range match {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		request := engine.Request{
			Url: string(m[1]),
			ParseFunc: func(contents []byte) engine.ParseResult {
				return ParsePerson(contents, name)
			},
		}
		result.Requests = append(result.Requests, request)
	}
	return result
}
