package main

//#include <stdio.h>
import "C"
import (
	MyBuffer "cgo/object_from_c++_2_go/include"
	"unsafe"
)

func main() {
	buf := MyBuffer.NewMyBuffer(1024)
	defer buf.Delete()

	copy(buf.Data(), []byte("hello\x00"))
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
