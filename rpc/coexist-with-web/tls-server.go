package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	port := ":1011"
	http.ListenAndServeTLS(port, "E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\server.crt", "E:\\Knows\\Golang\\GoPath\\src\\rpc\\Protobuf\\gRPC\\tls\\server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mux.ServeHTTP(w, r)
			return
		}),
	)
}
