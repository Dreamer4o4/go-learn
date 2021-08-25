package test

import (
	"golearn/src/cache"
	"strconv"
	"testing"
)

func TestCachePool(t *testing.T) {
	cm := cache.NewCachePool(":4000", func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})
	go cm.Run(cache.GetCache(func(key string) ([]byte, bool) {
		t.Log("try to find : ", key)
		return nil, true
	}))

	cm.AddRmoteCache("1", 1)
	cm.AddRmoteCache("2", 1)
	cm.AddRmoteCache("4", 2)

	t.Log(cm.GetPeer("3"))
	t.Log(cm.GetPeer("23"))
	t.Log(cm.GetPeer("40"))
	t.Log(cm.GetPeer("41"))
	t.Log(cm.GetPeer("42"))
}

func TestCacheClient(t *testing.T) {
	cm := cache.NewCachePool(":4001", func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})
	go cm.Run(cache.GetCache(func(key string) ([]byte, bool) {
		t.Log("try to find : ", key)
		return nil, true
	}))

	cm.AddRmoteCache("127.0.0.1:4001", 1)
	ret, _ := cm.GetPeer("3").GetValue("3")
	t.Log(ret)
	ret, _ = cm.GetPeer("4").GetValue("4")
	t.Log(ret)
}
