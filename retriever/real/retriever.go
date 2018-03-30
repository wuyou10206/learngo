package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, error := http.Get(url)
	if error != nil {
		panic(error)
	}
	result, error := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if error != nil {
		panic(error)
	}
	return string(result)
}
