package main

import (
	"contribution/in-memory-cache/v1-http/api"
	"contribution/in-memory-cache/v1-http/cache"
)

func main() {
	c := cache.New("inmemory")
	api.New(c).Listen()
}
