package parse

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityListRe)
	match := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, m := range match {
		result.Items = append(result.Items, "City "+string(m[2]))
		request := engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		}
		result.Requests = append(result.Requests, request)
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
