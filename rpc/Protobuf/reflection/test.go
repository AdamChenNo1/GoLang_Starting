package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterYourOwnServer(s, &server{})

	reflection.Register(s)

	s.Serve(lis)
}
