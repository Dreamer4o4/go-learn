package cacheStrategy

type cacheManagement struct {
	maxCap  int64
	curSize int64
}

type cacheStrategy interface {
	Push(key string, value Value)
	Pop()
	Find(key string) Value
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
