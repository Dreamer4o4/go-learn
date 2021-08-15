package cacheStrategy

import "container/list"

type lruStrategy struct {
	ll *list.List
}
