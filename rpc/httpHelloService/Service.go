package main

import (
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	. "rpc/interface"
)

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal("ListenTCP error", err)
	}
}
