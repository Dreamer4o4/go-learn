package cacheStrategy

type cacheManagement struct {
	maxCap  int64
	curSize int64
}

type cacheStrategy interface {
	Push(key string, value Value)
	Pop(key string) Value
	Find(key string) Value
}

type Value interface {
	Size() int
}

type ByteValue struct {
	value []byte
}

func NewByteValue(str string) *ByteValue {
	return &ByteValue{
		value: []byte(str),
	}
}

func (bv *ByteValue) Size() int {
	return len(bv.value)
}
