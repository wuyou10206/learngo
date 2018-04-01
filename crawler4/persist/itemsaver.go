package persist

import (
	"context"
	"errors"
	"github.com/gpmgo/gopm/modules/log"
	"gopkg.in/olivere/elastic.v5"
	"learngo/crawler4/engine"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Warn("Item Server:got item #%d:%v", itemCount, item)
			itemCount++
			err := save(item)
			if err != nil {
				log.Warn("Item Saver error: save item %v:%v", item, err)
			}

		}

	}()
	return out
}
func save(item engine.Item) error {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //内网无法sniff
	if err != nil {
		return err
	}
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().Index("dating_profile").Type(item.Type)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func save2(item engine.Item) (string, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //内网无法sniff
	if err != nil {
		return "", err
	}
	if item.Type == "" {
		return "", errors.New("must supply Type")
	}
	indexService := client.Index().Index("dating_profile").Type(item.Type)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	resp, err := indexService.BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
