package main

import (
	"cache_v3/api/http"
	"cache_v3/api/tcp"
	"cache_v3/cache"
	"flag"
	"log"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is", *typ)
	ca := cache.New(*typ)
	go tcp.New(ca).Listen()
	http.New(ca).Listen()
}
