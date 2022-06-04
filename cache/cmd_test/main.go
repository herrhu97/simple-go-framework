package main

import (
	"fmt"
	"github.com/herrhu97/simple-go-framework/cache"

	"github.com/herrhu97/simple-go-framework/log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	cache.NewGroup("scores", 2<<10, cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Debug("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := cache.NewHTTPPool(addr)
	log.Debug("cache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
