package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 6 * time.Second
	endpoints      = []string{"https://35.234.29.146:2379"}
)

func main() {
	var etcdCa = "E:/Knows/Golang/GoPath/src/contribution/etcd/tls/ca.pem"
	var etcdCert = "E:/Knows/Golang/GoPath/src/contribution/etcd/tls/client.pem"
	var etcdCertKey = "E:/Knows/Golang/GoPath/src/contribution/etcd/tls/client-key.pem"

	cert, err := tls.LoadX509KeyPair(etcdCert, etcdCertKey)
	if err != nil {
		log.Fatal("err")
	}
	ca, err := ioutil.ReadFile(etcdCa)
	if err != nil {
		log.Fatal("err")
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	_tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		TLS:       _tlsConfig,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer cli.Close()

	key1, value1 := "testkey1", "value"

	//kv := clientv3.NewKV(cli)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Put(ctx, key1, value1)
	// if err != nil {
	// 	panic(err)
	// }
	cancel()
	if err != nil {
		log.Println("Put failed. ", err)
	} else {
		log.Printf("Put {%s:%s} succeed\n", key1, value1)
	}
}
