package persist

import "github.com/gpmgo/gopm/modules/log"

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Warn("Item Server:got item #%d:%v", itemCount, item)
			itemCount++
			save(item)
		}

	}()
	return out
}
func save(Item interface{}) {

}
