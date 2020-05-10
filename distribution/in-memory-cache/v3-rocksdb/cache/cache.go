package cache

import (
	"log"
	"sync"
)

type Cache interface {
	Set(string, []byte) error   //将键值对设置进缓存，当Set操作成功时，error返回 nil
	Get(string) ([]byte, error) //根据 key从缓存中获取 value,接收string类型的参数,返回［］byte和error
	Del(string) error           //从缓存中删除key，
	GetStat() Stat              //用于获取缓存的状态
}

/*
*New 创建并返回一个 Cache 接口
 */
func New(typ string) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache()
	}
	if typ == "rocksdb" {
		c = newRocksdbCache()
	}
	if c == nil {
		panic("Unknown cache type:" + typ)
	}
	log.Println(typ, "ready to save")
	return c
}

type inMemoryCache struct {
	c     map[string][]byte //用来保存键值对
	mutex sync.RWMutex      //对map的并发访问提供读写锁保护
	Stat                    //用来记录缓存状态
}

//Set
func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	tmp, exist := c.c[k]
	if exist {
		c.del(k, tmp)
	}
	c.c[k] = v
	c.add(k, v)
	return nil
}

//Get
func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.c[k], nil
}

//Del
func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, v)
	}
	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}

func newInMemoryCache() *inMemoryCache {
	return &inMemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}
