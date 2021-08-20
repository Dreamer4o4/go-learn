package cacheStrategy

import (
	"container/list"
	"fmt"
)

type lruStrategy struct {
	lruList *list.List
	lruMap  map[string]*list.Element
	cacheSize
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
		cacheSize: cacheSize{
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

	for cache.curSize > cache.maxCap {
		cache.popOldElement()
	}
}

func (cache *lruStrategy) Pop(key string) (Value, bool) {
	if listElement, ok := cache.lruMap[key]; ok {
		element := listElement.Value.(*lruElement)
		cache.curSize -= int64(element.Size())
		cache.lruList.Remove(listElement)
		delete(cache.lruMap, element.key)
		return element.value, ok
	}
	return nil, false
}

func (cache *lruStrategy) Find(key string) (Value, bool) {
	if listElement, ok := cache.lruMap[key]; ok {
		cache.lruList.MoveToFront(listElement)
		element := listElement.Value.(*lruElement)
		return element.value, ok
	}
	return nil, false
}

func (cache *lruStrategy) popOldElement() {
	if cache.lruList.Len() != 0 {
		oldListElement := cache.lruList.Back()
		element := oldListElement.Value.(*lruElement)
		cache.Pop(element.key)
	}
}

func (cache *lruStrategy) Show() {
	fmt.Println("---------")
	for cur := cache.lruList.Front(); cur != nil; cur = cur.Next() {
		fmt.Print(cur.Value.(*lruElement).key)
		fmt.Println(cur.Value.(*lruElement).value)
	}
	fmt.Println("---------")
}
