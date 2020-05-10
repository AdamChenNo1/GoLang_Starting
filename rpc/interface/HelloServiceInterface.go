package HelloService

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "HelloService"

type HelloService struct {
	Conn    net.Conn
	IsLogin bool
}

func (p *HelloService) Login(request string, reply *string) error {
	if request != "user:password" {
		return fmt.Errorf("auth:failed")
	}

	log.Println("login ok")
	p.IsLogin = true
	return nil
}

func (p *HelloService) Hello(request string, reply *string) error {
	if !p.IsLogin {
		return fmt.Errorf("please login")
	}

	*reply = "hello:" + request + ",from" + p.Conn.RemoteAddr().String()
	return nil
}

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type HelloServiceClient struct {
	*rpc.Client
}

var _HelloServiceInterface = (*HelloServiceClient)(nil)

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
