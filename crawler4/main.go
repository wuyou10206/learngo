package main

import (
	"learngo/crawler4/engine"
	"learngo/crawler4/persist"
	"learngo/crawler4/scheduler"
	"learngo/crawler4/zhenai/parse"
)

func main() {
	request := engine.Request{
		//	Url:       "http://www.zhenai.com/zhenghun",
		//ParseFunc: parse.ParseCityList,
		Url:       "http://www.zhenai.com/zhenghun/yulin",
		ParseFunc: parse.ParseCity,
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
		ItemChan:  persist.ItemSaver(),
	}
	e.Run(request)
	//resp, err := http.Get("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error:status code:", resp.StatusCode)
	//	return
	//}
	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)
	//if err != nil {
	//	panic(err)
	//}
	//printCity(all)
	//	fmt.Printf("%s\n", all)
}

//func printCity(contents []byte) {
//	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`)
//	match := reg.FindAllSubmatch(contents, -1)
//	for _, m := range match {
//		for _, subMatch := range m[1:] {
//			fmt.Printf("%s\t", subMatch)
//		}
//		fmt.Println()
//	}
//	//	fmt.Printf("%s\n", match)
//}
//func determineEncoding(r io.Reader) encoding.Encoding {
//
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		panic(err)
//	}
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//	return e
//}
