package watch

import (
	"fmt"
	"sync"
)

//KVStoreService 内存数据库
type KVStoreService struct {
	m      map[string]string           //用于存储KV数据
	filter map[string]func(key string) //对应每个Watch调用时定义的过滤器函数列表
	mu     sync.Mutex                  //互斥锁，用于在多个Goroutine访问或修改时对其它成员提供保护
}

// NewKVStoreService 新建数据库
func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

//Get 取值函数
func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

//Set 存值函数
func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}
