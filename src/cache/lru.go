package cache

type lruStrategy struct {
	ll *list.List
}
