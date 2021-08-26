package cache

import (
	cacheStrategy "golearn/src/cache/cachestrategy"
	"log"
	"sync"
)

type CacheMangement struct {
	mtx        sync.Mutex
	localCache cacheStrategy.CacheStrategy
	peer       cacheServer
	loader     loadCallBack
}

type cacheServer interface {
	GetValue(key string) ([]byte, error)
}

type loadCallBack interface {
	Load(key string) ([]byte, bool)
}

type LoadFunc func(key string) ([]byte, bool)

func (lf LoadFunc) Load(key string) ([]byte, bool) {
	return lf(key)
}

func NewCacheMangement(local cacheStrategy.CacheStrategy, peer cacheServer, load loadCallBack) *CacheMangement {
	return &CacheMangement{
		localCache: local,
		peer:       peer,
		loader:     load,
	}
}

func (cm *CacheMangement) GetValue(key string) ([]byte, bool) {
	res, err := cm.peer.GetValue(key)
	if err != nil {
		log.Fatal(err)
		return nil, false
	}
	return res, true
}

func (cm *CacheMangement) GetValueLocal(key string) ([]byte, bool) {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()

	if value, ok := cm.localCache.Find(key); ok {
		return value.GetByteValue(), ok
	}

	if value, ok := cm.loader.Load(key); ok {
		cm.localCache.Push(key, cacheStrategy.NewByteValue(string(value)))
		return value, ok
	}

	return nil, false
}
