package cache

import (
	"errors"
	httpServer "golearn/src/httpserver"
	"io/ioutil"
	"net/http"
	"sync"
)

type CachePool struct {
	mtx        sync.Mutex
	hashmap    *ConsistentHash
	localcache *httpServer.HttpServer
}

type cacheClient struct {
	Protocol  string
	PeerAddr  string
	CachePath string
}

const defultProtocol string = "http"
const defultCachePath string = "/cache/"

func NewCachePool(localAddr string, hash Hash) *CachePool {
	return &CachePool{
		hashmap:    NewConsistentHash(hash),
		localcache: httpServer.NewHttpServer(localAddr),
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

func (cp *CachePool) init() {
	cp.localcache.AddGetGroupFunc(defultCachePath, func(ctxt *httpServer.Context) {

	})
}

func (cp *CachePool) run() {
	cp.init()
	cp.localcache.Run()
}

func newCacheClient(peerAddr, protocol, cachePath string) *cacheClient {
	return &cacheClient{
		PeerAddr:  peerAddr,
		Protocol:  protocol,
		CachePath: cachePath,
	}
}

func (cl *cacheClient) GetValue(key string) ([]byte, error) {
	url := cl.Protocol + "://" + cl.PeerAddr + cl.CachePath + key
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad request : " + string(resp.Status))
	}

	value, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return value, nil
}
