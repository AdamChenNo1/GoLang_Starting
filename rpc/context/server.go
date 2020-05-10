package main

import (
	"log"
	"net"
	"net/rpc"
	hs "rpc/interface"
)

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP Error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept Error:", err)
		}

		go func() {
			defer conn.Close()

			p := rpc.NewServer()
			p.Register(&hs.HelloService{Conn: conn, IsLogin: false})
			p.ServeConn(conn)
		}()
	}
}
