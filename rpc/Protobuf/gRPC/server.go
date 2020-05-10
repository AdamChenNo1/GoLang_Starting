package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	ft "rpc/Protobuf/gRPC/filter"
	hs "rpc/Protobuf/gRPC/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	// creds, err := credentials.NewServerTLSFromFile("E:/Knows/Golang/GoPath/src/rpc/Protobuf/gRPC/tls/server.crt", "E:/Knows/Golang/GoPath/src/rpc/Protobuf/gRPC/tls/server.key") //从文件为服务器构造证书对象
	// if err != nil {
	// 	log.Fatal(err)
	// }

	certificate, err := tls.LoadX509KeyPair("E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\server.crt", "E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\server.key")

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(ft.Filter)) //证书包装为选项后作为参数传入grpc.NewServer函数
	hs.RegisterHelloServiceServer(grpcServer, new(hs.HelloServiceImpl))
	if err != nil {
		log.Fatal(err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Serve(lis)
}
