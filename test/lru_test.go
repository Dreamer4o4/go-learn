package test

import (
	cacheStrategy "golearn/src/cache/cacheStrategy"
	"testing"
)

func TestLru(t *testing.T) {
	lru := cacheStrategy.NewLruStrategy(10)
	lru.Push("1", cacheStrategy.NewByteValue("asd"))
	lru.Push("2", cacheStrategy.NewByteValue("a"))
	lru.Push("3", cacheStrategy.NewByteValue("zx"))
	lru.Show()
	lru.Push("4", cacheStrategy.NewByteValue("qwe"))
	lru.Show()

	if v, ok := lru.Find("2"); ok {
		t.Log(v)
	}
	if v, ok := lru.Find("1"); ok {
		t.Log(v)
	}
	lru.Show()

}
