package cacheStrategy

import (
	"container/list"
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

func NewLruStrategy(maxSize int64) *lruStrategy {
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
	if listElement, ok := cache.lruMap[key]; ok {
		cache.lruList.MoveToFront(listElement)
		element := listElement.Value.(*lruElement)
		cache.curSize += (int64(value.Size()) - int64(element.value.Size()))
		element.value = value
	} else {
		listElement := cache.lruList.PushFront(&lruElement{key, value})
		element := listElement.Value.(*lruElement)
		cache.lruMap[key] = listElement
		cache.curSize += int64(element.Size())
	}
}

func (cache *lruStrategy) Pop(key string) Value {
	if listElement, ok := cache.lruMap[key]; ok {

	}
}
