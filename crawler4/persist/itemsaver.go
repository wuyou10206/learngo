package persist

import (
	"context"
	"github.com/gpmgo/gopm/modules/log"
	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Warn("Item Server:got item #%d:%v", itemCount, item)
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Warn("Item Saver error: save item %v:%v", item, err)
			}

		}

	}()
	return out
}
func save(item interface{}) (string, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //内网无法sniff
	if err != nil {
		return "", err
	}
	resp, err := client.Index().Index("dating_profile").Type("zhenai").
		BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
