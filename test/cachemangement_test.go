package test

import (
	"golearn/src/cache"
	cacheStrategy "golearn/src/cache/cachestrategy"
	"testing"
)

func TestCacheMangement(t *testing.T) {
	db := map[string]string{
		"1": "111",
		"2": "222",
		"3": "333",
	}

	cp := cache.NewCachePool(":4000", nil)
	cm := cache.NewCacheMangement(cacheStrategy.NewLruStrategy(100), cp, cache.LoadFunc(func(key string) ([]byte, bool) {
		if v, ok := db[key]; ok {
			t.Log("load : ", key)
			return []byte(v), ok
		}
		t.Log("query nil : ", key)
		return nil, false
	}))
	go cp.Run(cache.GetCache(cm.GetValueLocal))
	cp.AddRmoteCache("127.0.0.1:4000", 1)

	cm.GetValue("1")
	cm.GetValue("2")
	cm.GetValue("1")
	cm.GetValue("3")
	cm.GetValue("1")

}
