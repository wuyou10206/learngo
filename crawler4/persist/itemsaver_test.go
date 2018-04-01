package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"learngo/crawler4/model"
	"testing"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
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
	id, err := save(profile)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("dating_profile").Type("zhenai").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	if profile != actual {
		t.Errorf("got %v ; expected %v", actual, profile)
	}
}
