package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	. "rpc/interface"
)

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
