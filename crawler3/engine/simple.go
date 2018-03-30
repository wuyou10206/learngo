package engine

import (
	"learngo/crawler3/fetcher"

	"github.com/gpmgo/gopm/modules/log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Warn("Got item %v", item)
		}
	}
}
func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	log.Warn("Fetching %v", r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
