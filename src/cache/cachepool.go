package cache

import (
	"net/http"
	"sync"
)

type CachePool struct {
	mtx     sync.Mutex
	hashmap ConsistentHash
}

type cacheClient struct {
	Protocol  string
	PeerAddr  string
	CachePath string
}

const defultProtocol string = "http"
const defultCachePath string = "cache/"

func NewCachePool(hash Hash) *CachePool {
	return &CachePool{
		hashmap: *NewConsistentHash(hash),
	}
}

func (cp *CachePool) AddRmoteCache(addr string, virtualNodeNum int) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	cp.hashmap.AddRealServer(*NewRealServerNode(addr, virtualNodeNum))
}

func (cp *CachePool) RemoveRemoteCache(addr string) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	cp.hashmap.RemoveRealServer(addr)
}

func (cp *CachePool) GetPeer(key string) peerCache {
	if peerAddr := cp.hashmap.FindServer(key); peerAddr != "" {
		peerClient := newCacheClient(peerAddr, defultProtocol, defultCachePath)
		return peerClient
	}
	return nil
}

func newCacheClient(peerAddr, protocol, cachePath string) *cacheClient {
	return &cacheClient{
		PeerAddr:  peerAddr,
		Protocol:  protocol,
		CachePath: cachePath,
	}
}

func (cl *cacheClient) GetValue(key string) []byte {
	url := cl.Protocol + "://" + cl.PeerAddr + "/" + cl.CachePath + key
	_, err := http.Get(url)
	if err != nil {
		return nil
	}
	return nil
}
