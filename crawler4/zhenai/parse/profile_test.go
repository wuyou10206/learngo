package parse

import (
	"io/ioutil"
	"learngo/crawler/model"
	"learngo/crawler4/engine"
	"testing"
)

func TestParsePerson(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("test_profile_data.html")
	//fmt.Printf("%s", contents)
	if err != nil {
		panic(err)
	}
	result := ParsePerson(contents, "", "风中的蒲公英照片")

	if len(result.Items) != 1 {
		t.Errorf("实际 %d ,期望 %d", len(result.Items), 1)
	}
	profile := result.Items[0]
	expected := model.Profile{
		Name:       "风中的蒲公英照片",
		Gender:     "女",
		Age:        41,
		Height:     158,
		Weight:     48,
		Income:     "3001-5000元",
		Marriage:   "离异",
		Education:  "中专",
		Occupation: "公务员",
		Hokou:      "四川阿坝",
		Xingzuo:    "处女座",
		House:      "已购房",
		Car:        "未购车",
	}
	item := engine.Item{
		Url:     "http://1111111",
		Type:    "zhenai",
		Id:      "112233444",
		Payload: expected,
	}
	if item != profile {
		t.Errorf("实际 %v ,期望 %v", profile, expected)
	}

}
