package main

import (
	"log"
	"net/http"
	pb_hello "rpc/Protobuf/pbgo/service"
)

func main() {
	router := pb_hello.HelloServiceHandler(new(pb_hello.HelloService))
	log.Fatal(http.ListenAndServe(":8080", router))
}
