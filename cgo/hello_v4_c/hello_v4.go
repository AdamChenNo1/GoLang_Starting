//hello_v4.go
package main

//#include "hello.h"
import "C"

func main() {
	C.SayHello(C.CString("Hello World\n"))
}