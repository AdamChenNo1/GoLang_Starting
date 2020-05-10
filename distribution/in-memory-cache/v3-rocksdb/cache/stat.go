package cache

type Stat struct {
	Count     int64 //表示缓存目前保存的键值对数量
	KeySize   int64 //key占据的总字节数
	ValueSize int64 //value占据的总字节数
}

//add 用于新增键值对时改变缓存的状态
func (s *Stat) add(k string, v []byte) {
	s.Count += 1
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

//del 用于删除键值对时改变缓存的状态
func (s *Stat) del(k string, v []byte) {
	s.Count -= 1
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
