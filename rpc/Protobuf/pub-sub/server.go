package main

import (
	"context"
	"log"
	hs "rpc/Protobuf/pub-sub/cross"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hs.NewPubsubServiceClient(conn)

	_, err = client.Publish(context.Background(), &hs.String{Value: "golang:hello Go"})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Publish(context.Background(), &hs.String{Value: "docker:hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}
