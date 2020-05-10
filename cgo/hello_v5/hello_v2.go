//hello_v5.go
package main

//void SayHello(_GoString_ s);
import "C"

import "fmt"

func main() {
	C.SayHello("Hello World")
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}
