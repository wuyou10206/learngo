package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
	"time"
)

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) string {
	return poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
	//	Connect(host string)
}

func session(s RetrieverPoster) string {
	//s.Get("")
	s.Post(url, map[string]string{
		"contents": "大家好",
	})
	return s.Get(url)
}
func main() {
	var r Retriever
	retriever := mock.Retriever{"besttest"}
	//	fmt.Printf("%T %v\n", r, r)
	r = &retriever
	inspect(r)
	fmt.Println(download(r))
	fmt.Println(download(&mock.Retriever{"BestTest"}))
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)
	mockRetriever, ok := r.(*mock.Retriever)
	if ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Println("try session")
	fmt.Println(session(&retriever))
	//fmt.Println(post(retriever))
	//	fmt.Printf("%T %v\n", r, r)
	//	fmt.Println(download(r))

}
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:" + v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:" + v.UserAgent)
	}
	fmt.Println()
}
