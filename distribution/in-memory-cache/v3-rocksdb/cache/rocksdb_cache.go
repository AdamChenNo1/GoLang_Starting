package cache

//#include "rocksdb/c.h"
//#cgo CFLAGS: -I${SRCDIR}/../../../rocksdb/include
//#cgo LDFLAGS:-L${SRCDIR}/../../../rocksdb -lrocksdb -lz -lphread -lsnappy -lstdc++ -lm -O3
import "C"

import (
	"errors"
	"log"
	"regexp"
	"runtime"
	"strconv"
	"unsafe"
)

type rocksdbCache struct {
	db *C.rocksdb_t              //表示一个RocksDB 存储的实例
	ro *C.rocksdb_readoptions_t  //读操作的选项类型
	wo *C.rocksdb_writeoptions_t //写操作的选项类型
	e  *C.char                   //用来指向 RocksDB C API 返回的错误字符串
}

func newRocksdbCache() *rocksdbCache {
	options := C.rocksdb_options_create()                                    //创建一个rocksdb options_t 类型的指针 options
	C.rocksdb_options_increase_parallelism(options, C.int(runtime.NumCPU())) //中设置 RocksDB 并发线程数
	C.rocksdb_options_set_create_if_missing(options, 1)                      //如果目标目录不存在则创建一个新的存储目录
	var e *C.char
	db := C.rocksdb_open(options, C.CString("/mnt/rocksdb"), &e) //打开位于 /mnt/rocksdb/的存储目录
	if e != nil {
		log.Panicln(C.GoString(e))
	}
	C.rocksdb_options_destroy(options) //销毁options
	return &rocksdbCache{
		db,
		C.rocksdb_readoptions_create(),
		C.rocksdb_writeoptions_create(),
		e,
	}
}

func (c *rocksdbCache) Get(key string) ([]byte, error) {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k)) //k会在整个服务进程的生命周期内被多次调用,所以需要在函数退出时用free释放k的内存
	var length C.size_t
	v := C.rocksdb_get(c.db, c.ro, k, C.size_t(len(key)), &length, &c.e)
	if c.e != nil {
		return nil, errors.New(C.GoString(c.e))
	}
	defer C.free(unsafe.Pointer(v))
	return C.GoBytes(unsafe.Pointer(v), C.int(length)), nil
}

func (c *rocksdbCache) Set(key string, value []byte) error {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k)) //k会在整个服务进程的生命周期内被多次调用,所以需要在函数退出时用free释放k的内存
	v := C.CBytes(value)
	defer C.free(unsafe.Pointer(v))
	C.rocksdb_put(c.db, c.wo, k, C.size_t(len(key)), (*C.char)(v), C.size_t(len(value)), &c.e)
	if c.e != nil {
		return errors.New(C.GoString(c.e))
	}
	return nil
}

func (c *rocksdbCache) Del(key string) error {
	k := C.CString(key)
	defer C.free(unsafe.Pointer(k)) //k会在整个服务进程的生命周期内被多次调用,所以需要在函数退出时用free释放k的内存
	C.rocksdb_delete(c.db, c.wo, k, C.size_t(len(key)), &c.e)
	if c.e != nil {
		return errors.New(C.GoString(c.e))
	}
	return nil
}

func (c *rocksdbCache) GetStat() Stat {
	k := C.CString("rocksdb.aggregated-table-propertites")
	defer C.free(unsafe.Pointer(k))
	v := C.rocksdb_property_value(c.db, k)
	defer C.free(unsafe.Pointer(v))
	p := C.GoString(v)
	r := regexp.MustCompile(`([^;]+)=([^;]+);`)
	s := Stat{}
	for _, submatches := range r.FindAllStringSubmatch(p, -1) {
		if submatches[1] == "# entries" {
			s.Count, _ = strconv.ParseInt(submatches[2], 10, 64)
		} else if submatches[1] == " raw key size" {
			s.KeySize, _ = strconv.ParseInt(submatches[2], 10, 64)
		} else if submatches[1] == " raw value size" {
			s.ValueSize, _ = strconv.ParseInt(submatches[2], 10, 64)
		}
	}
	return s

}
