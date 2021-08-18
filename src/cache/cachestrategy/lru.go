package cacheStrategy

import (
	"container/list"
	"fmt"
)

type lruStrategy struct {
	lruList *list.List
	lruMap  map[string]*list.Element
	cacheManagement
}

type lruElement struct {
	key   string
	value Value
}

func (le *lruElement) Size() int {
	return len(le.key) + le.value.Size()
}

func newLruStrategy(maxSize int64) *lruStrategy {
	return &lruStrategy{
		cacheManagement: cacheManagement{
			maxCap:  maxSize,
			curSize: 0,
		},
		lruList: list.New(),
		lruMap:  make(map[string]*list.Element),
	}
}

func (cache *lruStrategy) Push(key string, value Value) {
	if v, ok := cache.lruMap[key]; ok {
		cache.lruList.MoveToFront(v)
		element := v.Value.(*lruElement)
		fmt.Print(element.Size())
	} else {
		element := cache.lruList.PushFront(&lruElement{key, value})
		cache.lruMap[key] = element
	}
}
