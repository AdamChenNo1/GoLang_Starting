package main

import (
	"context"
	"fmt"
	"io"
	"log"

	hs "rpc/Protobuf/pub-sub/cross"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(
		"tls/server.crt", "server.grpc.io",
	)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hs.NewPubSubServiceClient(conn)
	stream, err := client.Subscribe(context.Background(), &hs.String{Value: "golang"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}
}
