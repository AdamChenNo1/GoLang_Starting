package main

//#cgo LDFLAGS:-LE:/Knows/Golang/GoPath/src/cgo/number -lnumber
//
//#include "number/include/number.h"
import "C"
import "fmt"

func main() {
	fmt.Println(C.number_add_mod(10, 5, 12))
}
