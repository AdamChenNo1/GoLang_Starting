package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/zieckey/etcdsync"
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

	m, err := etcdsync.New("/lock", 10, []string{"http://35.221.227.172:2379"})
	if m == nil || err != nil {
		log.Printf("etcdsync.New failed")
		return
	}

	err = m.Lock()
	if err != nil {
		log.Printf("etcdsync.Lock failed")
		return
	}

	log.Printf("etcdsync.Lock OK")
	log.Printf("Get the lock.Do something here")

	err = m.Unlock()
	if err != nil {
		log.Printf("etcdsync.Unlock failed")
	} else {
		log.Printf("etcdsync.Unlock OK")
	}

}
