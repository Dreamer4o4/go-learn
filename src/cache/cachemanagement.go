package cache

type cacheManagement struct {
	maxCap  int64
	curSize int64
}

type cacheStrategy interface {
	push(key string, value Value)
	pop()
	update(key string, value Value)
	find(key string) Value
}

type Value interface {
	Size() int
}

type ByteValue struct {
	value []byte
}

func (bv *ByteValue) Size() int {
	return len(bv.value)
}
