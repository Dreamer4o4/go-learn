package cache

import (
	cacheStrategy "golearn/src/cache/cachestrategy"
	"sync"
)

type CacheMangement struct {
	mtx        sync.Mutex
	localCache cacheStrategy.CacheStrategy
	picker     peerPicker
	load       loadCallBack
}

type peerPicker interface {
	GetPeer(key string) peerCache
}

type peerCache interface {
	GetValue(key string) ([]byte, error)
}

type loadCallBack interface {
	Load(key string) []byte
}

func NewCacheMangement(local cacheStrategy.CacheStrategy, picker peerPicker, load loadCallBack) *CacheMangement {
	return &CacheMangement{
		localCache: local,
		picker:     picker,
		load:       load,
	}
}

func (cm *CacheMangement) GetValue(key string) ([]byte, bool) {
	return cm.getValueLocal(key)
}

func (cm *CacheMangement) getValueLocal(key string) ([]byte, bool) {
	cm.mtx.Lock()
	defer cm.mtx.Unlock()

	if value, ok := cm.localCache.Find(key); ok {
		return value.GetByteValue(), ok
	}

	return nil, false
}
