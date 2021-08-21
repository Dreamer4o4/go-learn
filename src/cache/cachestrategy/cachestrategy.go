package cacheStrategy

import (
	"bytes"
	"encoding/gob"
)

type cacheSize struct {
	maxCap  int64
	curSize int64
}

type CacheStrategy interface {
	Push(key string, value Value)
	Pop(key string) (Value, bool)
	Find(key string) (Value, bool)
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

func ToBytes(value interface{}) ([]byte, error) {
	var res bytes.Buffer
	enc := gob.NewEncoder(&res)
	if err := enc.Encode(value); err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}
