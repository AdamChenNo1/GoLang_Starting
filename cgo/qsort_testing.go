package main

//extern int go_qsort_compare(void* a,void* b);
import "C"

import (
	qsort_v3 "cgo/qsort/v3"
	"fmt"
	"unsafe"
)

//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

func main() {
	values := []int64{42, 9, 8, 7, 2, 6, 57}

	// qsort_v1.Sort(unsafe.Pointer(&values[0]), len(values), int(unsafe.Sizeof(values[0])), qsort.CompareFunc(C.go_qsort_compare))
	// qsort_v2.Sort_v2(
	// 	unsafe.Pointer(&values[0]),
	// 	len(values),
	// 	int(unsafe.Sizeof(values[0])),
	// 	func(a, b unsafe.Pointer) int {
	// 		pa, pb := (*int32)(a), (*int32)(b)
	// 		return int(*pa - *pb)
	// 	},
	// )
	qsort_v3.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values)
}
