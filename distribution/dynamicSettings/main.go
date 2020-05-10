package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"go.etcd.io/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 6 * time.Second
	endpoints      = []string{"https://35.234.29.146:2379","https://35.234.29.146:22379","https://35.234.29.146:32379"}
)

type ConfigStruct struct {
	Addr string `json:"addr"`
	AesKey string `json:"aes_key"`
	HTTPS bool `json:"https"`
	Secret string `json:"secret"`
	PrivateKeyPath string `json:"private_key_path"`
	CertFilePath string `json:"cert_file_path"`
}

var appConfig ConfigStruct

func init(){
	var etcdCa = "E:/Knows/Golang/GoPath/src/contribution/etcd/tls/ca.pem"
	var etcdCert = "E:/Knows/Golang/GoPath/src/contribution/etcd/tls/client.pem"
	var etcdCertKey = "E:/Knows/Golang/GoPath/src/contribution/etcd/tls/client-key.pem"

	cfg := client.Config{
		Endpoints:endpoints,
		TLS:_tlsConfig,
	}

	cert,err := tls.LoadX509KeyPair(etcdCert,etcdCertKey)
	if err != nil {
		log.Fatal("loading key pair failed",err)
	}

	ca,err := ioutil.ReadFile(etcdCa)
	if err != nil {
		log.Fatal("loading CA failed",err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	_tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs: certPool,
	}

	cli,err := clientv3.New(cfg)
	if err != nil {
		log.Fatal("creating clientv3 failed")
	}

	defer cli.Close()

}

func WatchAndUpdate(){
	
}

func initConfig(){
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp,err := cli.Get(ctx,configPath)
	if err != nil {
		log.Fatal("config getting failed")
	}

	err := json,Unmarshal(resp.Node.Value,&appConfig)
	if err != nil {
		log.Fatal("error unmarshalling failed")
	}
}

func getConfig(){
	return appConfig
}