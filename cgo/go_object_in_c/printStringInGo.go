package main

/*
extern char* NewGoString(char* );
extern void FreeGoString(char* );
extern void PrintGoString(char* );

static void printString(char* s) {
    char* gs = NewGoString(s);
    PrintGoString(gs);
    FreeGoString(gs);
}
*/
import "C"
import (
	printString "cgo/go_object_in_c/lib"
	"unsafe"
)

//export NewGoString
func NewGoString(s *C.char) *C.char {
	gs := C.GoString(s)
	id := printString.NewObjectId(gs)
	return (*C.char)(unsafe.Pointer(uintptr(id)))
}

//export FreeGoString
func FreeGoString(p *C.char) {
	id := printString.ObjectId(uintptr(unsafe.Pointer(p)))
	id.Free()
}

//export PrintGoString
func PrintGoString(p *C.char) {
	id := printString.ObjectId(uintptr(unsafe.Pointer(p)))
	gs := id.Get().(string)
	print(gs)
}

func main() {
	C.printString(C.CString("hello"))
	// C.printString("hello")
}
