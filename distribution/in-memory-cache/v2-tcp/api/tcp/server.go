package tcp

import (
	"bufio"
	"cache_v2/cache"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

//Server 负责处理TCP连接以及和客户端的交互
type Server struct {
	cache.Cache
}

//Listen 监听Tcp 12346端口请求
func (s *Server) Listen() {
	l, e := net.Listen("tcp", ":12346")
	if e != nil {
		log.Panicln(e)
	}
	for {
		c, e := l.Accept()
		if e != nil {
			log.Panicln(e)
		}
		go s.process(c)
	}
}

//New 新建cache Tcp Server
func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server) process(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for { //在一个无限循环中调用bufio.Reader.ReadByte方法获取command中的 op部分，并根据op决定调用get、set或del方法
		op, e := r.ReadByte()
		if e != nil {
			if e != io.EOF {
				log.Println("close connection due to error:", e)
			}
			return
		}
		if op == 'S' {
			e = s.set(conn, r)
		} else if op == 'G' {
			e = s.get(conn, r)
		} else if op == 'D' {
			e = s.del(conn, r)
		} else {
			log.Println("close connection due to invilad operation:", op)
			return
		}
		if e != nil {
			log.Println("close connection due to error:", e)
			return
		}
	}
}

func (s *Server) set(conn net.Conn, r *bufio.Reader) error {
	k, v, e := s.readKeyAndValue(r)
	if e != nil {
		return e
	}
	return sendResponse(nil, s.Set(k, v), conn)
}

func (s *Server) get(conn net.Conn, r *bufio.Reader) error {
	k, e := s.readKey(r)
	if e != nil {
		return e
	}
	v, e := s.Get(k)
	return sendResponse(v, e, conn)
}

func (s *Server) del(conn net.Conn, r *bufio.Reader) error {
	k, e := s.readKey(r)
	if e != nil {
		return e
	}
	return sendResponse(nil, s.Del(k), conn)
}

//sendResponse 将服务端的error或value写入客户端连接
func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		tmp := fmt.Sprintf("-%d", len(errString)) + errString
		_, e := conn.Write([]byte(tmp))
		return e
	}
	vlen := fmt.Sprintf("%d", len(value))
	_, e := conn.Write(append([]byte(vlen), value...))
	return e
}

//解析客户端发来的command,从中获取key
func (s *Server) readKey(r *bufio.Reader) (string, error) {
	klen, e := readLen(r)
	if e != nil {
		return "", e
	}
	k := make([]byte, klen)
	_, e = io.ReadFull(r, k)
	if e != nil {
		return "", e
	}
	return string(k), nil
}

//解析客户端发来的command，从中获取key和value
func (s *Server) readKeyAndValue(r *bufio.Reader) (string, []byte, error) {
	klen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}
	vlen, e := readLen(r)
	if e != nil {
		return "", nil, e
	}
	k := make([]byte, klen)
	_, e = io.ReadFull(r, k)
	if e != nil {
		return "", nil, e
	}
	v := make([]byte, vlen)
	_, e = io.ReadFull(r, v)
	if e != nil {
		return "", nil, e
	}
	return string(k), v, nil
}

//readLen 以空格为分隔符读取一个字符串并将之转化为一个整型
func readLen(r *bufio.Reader) (int, error) {
	tmp, e := r.ReadString(' ')
	if e != nil {
		return 0, e
	}
	l, e := strconv.Atoi(strings.TrimSpace(tmp))
	if e != nil {
		return 0, e
	}
	return l, nil
}
