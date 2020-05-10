package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	rs "rpc/Protobuf/grpc2rest/util"

	"google.golang.org/grpc"
)

func startWeb() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	err := rs.RegisterRestServiceHandlerFromEndpoint(ctx, mux, "localhost:5000", []grpc.DialOption{grpc.WithInsecure()})

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", mux)
}

func startRPC() {
	grpcServer := grpc.NewServer()
	rs.RegisterRestServiceServer(grpcServer, new(rs.RestServiceImpl))

	lis, _ := net.Listen("tcp", ":5000")
	grpcServer.Serve(lis)
}

func main() {
	go startWeb()
	startRPC()
}
