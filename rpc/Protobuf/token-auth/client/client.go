package main

import (
	"context"
	"fmt"
	"io"
	"log"
	. "rpc/Protobuf/token-auth/auth"
	. "rpc/Protobuf/token-auth/util"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// creds, err := credentials.NewClientTLSFromFile(
	// 	"E:/Knows/Golang/GoPath/src/rpc/Protobuf/gRPC/tls/server.crt", "server.io",
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// certificate, err := tls.LoadX509KeyPair("E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\client.crt", "E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\client.key")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// certPool := x509.NewCertPool()
	// ca, err := ioutil.ReadFile("E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\ca.crt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if ok := certPool.AppendCertsFromPEM(ca); !ok {
	// 	log.Fatal("failed to append ca certs")
	// }

	// creds := credentials.NewTLS(&tls.Config{
	// 	Certificates: []tls.Certificate{certificate},
	// 	ServerName:   "server.io", //服务器名称
	// 	RootCAs:      certPool,
	// })
	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//向服务端发送数据
	go func() {
		for {
			if err := stream.Send(&String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	//接收服务端返回的数据
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.GetValue())
	}
}
