package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	ft "rpc/Protobuf/gRPC/filter"
	hs "rpc/Protobuf/gRPC/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":1011"
)

func newgRPCServer() *grpc.Server {
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

	return grpcServer
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "hello")
	})
	return mux
}

func co_exist(mux *http.ServeMux, grpcServer *grpc.Server) {
	http.ListenAndServeTLS(port, "E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\server.crt", "E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor != 2 {
				mux.ServeHTTP(w, r)
				return
			}

			if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
				return
			}

			mux.ServeHTTP(w, r)
			return
		}),
	)
}
func main() {
	mux := newMux()
	grpcServer := newgRPCServer()
	co_exist(mux, grpcServer)
}
