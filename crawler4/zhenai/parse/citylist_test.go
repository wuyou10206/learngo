package parse

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	//fmt.Printf("%s", contents)
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	const resultSize = 470
	//	expectedCitys := []string{"City 阿坝", "City 阿克苏", "City 阿拉善盟"}
	expectedUrls := []string{"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng"}
	if len(result.Requests) != resultSize {
		t.Errorf("实际 %d ,期望 %d", len(result.Requests), resultSize)
	}
	//if len(result.Items) != resultSize {
	//	t.Errorf("实际 %d ,期望 %d", len(result.Items), resultSize)
	//}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("实际 %s ,期望 %s", result.Requests[i].Url, url)
		}
	}
	//for i, city := range expectedCitys {
	//	if result.Items[i].Payload.(string) != city {
	//		t.Errorf("实际 %s ,期望 %s", result.Items[i].Payload.(string), city)
	//	}
	//}
}
