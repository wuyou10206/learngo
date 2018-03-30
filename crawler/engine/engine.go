package engine

import (
	"learngo/crawler/fetcher"

	"github.com/gpmgo/gopm/modules/log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		log.Warn("Fetching %v", r.Url)
		if err != nil {
			log.Warn("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Warn("Got item %v", item)
		}
	}
}
