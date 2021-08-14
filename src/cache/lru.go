package cache

import "container/list"

type lruStrategy struct {
	lrulist *list.List
	lrumap  map[string]*list.Element
	cacheManagement
}
