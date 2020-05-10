package main

/*
#cgo CPPFLAGS:-IE://Knows/Golang/GoPath/src/cgo/object_from_go_2_c++/include
#include "person_capi.h"
*/
import "C"
import (
	objectid "cgo/go_object_in_c/lib"
	person "cgo/object_from_go_2_c++/src/person"
	"unsafe"
)

//export person_new
func person_new(name *C.char, age C.int) C.person_handle_t {
	id := objectid.NewObjectId(person.NewPerson(C.GoString(name), int(age)))
	return C.person_handle_t(id)
}

//export person_delete
func person_delete(p C.person_handle_t) {
	person := objectid.ObjectId(p)
	(&person).Free()
}

//export person_set
func person_set(p C.person_handle_t, name *C.char, age C.int) {
	h := objectid.ObjectId(p).Get().(*person.Person)
	h.Set(C.GoString(name), int(age))
}

//export person_get_name
func person_get_name(p C.person_handle_t, buf *C.char, size C.int) *C.char {
	h := objectid.ObjectId(p).Get().(*person.Person)
	name, _ := h.Get()

	n := int(size) - 1
	bufSlice := ((*[1 << 31]byte)(unsafe.Pointer(buf)))[0:n:n]
	n = copy(bufSlice, []byte(name))
	bufSlice[n] = 0
	return buf
}

//export person_get_age
func person_get_age(p C.person_handle_t) C.int {
	h := objectid.ObjectId(p).Get().(*person.Person)
	_, age := h.Get()
	return C.int(age)
}

func main() {}
