package test

import (
	"golearn/src/cache"
	cacheStrategy "golearn/src/cache/cachestrategy"
	"testing"
	"time"
)

func TestCacheMangement(t *testing.T) {
	db1 := map[string]string{
		"1": "111",
		"2": "222",
		"3": "333",
	}

	db2 := map[string]string{
		"12": "111",
		"22": "222",
		"32": "333",
	}

	cp := cache.NewCachePool(":4000", nil)
	cp.AddRmoteCache("127.0.0.1:4000", 10)
	cp.AddRmoteCache("127.0.0.1:4001", 10)
	cache.NewCacheMangement(cacheStrategy.NewLruStrategy(100), cp, cache.LoadFunc(func(key string) ([]byte, bool) {
		if v, ok := db1[key]; ok {
			t.Log("load : ", key)
			return []byte(v), ok
		}
		t.Log("query nil : ", key)
		return nil, false
	}))

	cp2 := cache.NewCachePool(":4001", nil)
	cp2.AddRmoteCache("127.0.0.1:4000", 10)
	cp2.AddRmoteCache("127.0.0.1:4001", 10)
	cache2 := cacheStrategy.NewLruStrategy(100)
	for k, v := range db2 {
		cache2.Push(k, cacheStrategy.NewByteValue(v))
	}
	cache.NewCacheMangement(cache2, cp2, cache.LoadFunc(func(key string) ([]byte, bool) {
		if v, ok := db2[key]; ok {
			t.Log("load : ", key)
			return []byte(v), ok
		}
		t.Log("query nil : ", key)
		return nil, false
	}))

	time.Sleep(1 * time.Second)
	client := cache.NewCachePool("", nil)
	client.AddRmoteCache("127.0.0.1:4000", 10)
	client.AddRmoteCache("127.0.0.1:4001", 10)
	client.GetValue("1")
	client.GetValue("2")
	client.GetValue("3")
	client.GetValue("12")
	client.GetValue("22")
	client.GetValue("32")
}
