package cache

import (
	"errors"
	httpServer "golearn/src/httpserver"
	"io/ioutil"
	"net/http"
	"sync"
)

type CachePool struct {
	mtx          sync.Mutex
	remoteServer *ConsistentHash
	localServer  *httpServer.HttpServer
	localCache   localCacheGetter
}

type cacheClient struct {
	Protocol  string
	PeerAddr  string
	CachePath string
}

type localCacheGetter interface {
	Query(key string) ([]byte, bool)
}

const defultProtocol string = "http"
const defultCachePath string = "/cache"

type GetCache func(key string) ([]byte, bool)

func (f GetCache) Query(key string) ([]byte, bool) {
	return f(key)
}

func NewCachePool(localAddr string, hash Hash) *CachePool {
	return &CachePool{
		remoteServer: NewConsistentHash(hash),
		localServer:  httpServer.NewHttpServer(localAddr),
	}
}

func (cp *CachePool) AddRmoteCache(addr string, virtualNodeNum int) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	cp.remoteServer.AddRealServer(*NewRealServerNode(addr, virtualNodeNum))
}

func (cp *CachePool) RemoveRemoteCache(addr string) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	cp.remoteServer.RemoveRealServer(addr)
}

func (cp *CachePool) GetPeer(key string) peerCache {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	if peerAddr := cp.remoteServer.FindServer(key); peerAddr != "" {
		peerClient := newCacheClient(peerAddr, defultProtocol, defultCachePath)
		return peerClient
	}
	return nil
}

/*
**	only for key query, basePath/key/nonsense
 */
func (cp *CachePool) init(basePath string) {
	cp.localServer.AddGetFunc(basePath+"/*", func(ctxt *httpServer.Context) {
		if paths := httpServer.ParasePath(ctxt.Req.URL.Path); len(paths) > 1 {
			queryKey := paths[len(httpServer.ParasePath(basePath))]

			if value, ok := cp.localCache.Query(queryKey); ok {
				//	cache hit
				ctxt.Resw.Write(value)
			} else {
				// 	cache miss
				ctxt.Resw.WriteHeader(http.StatusNotFound)
			}

			return
		}

		ctxt.Resw.WriteHeader(http.StatusBadRequest)
	})
}

func (cp *CachePool) registerCacheGetter(cb localCacheGetter) {
	cp.localCache = cb
}

func (cp *CachePool) Run(cb localCacheGetter) {
	cp.registerCacheGetter(cb)
	cp.init(defultCachePath)
	cp.localServer.Run()
}

func newCacheClient(peerAddr, protocol, cachePath string) *cacheClient {
	return &cacheClient{
		PeerAddr:  peerAddr,
		Protocol:  protocol,
		CachePath: cachePath,
	}
}

func (cl *cacheClient) GetValue(key string) ([]byte, error) {
	url := cl.Protocol + "://" + cl.PeerAddr + cl.CachePath + "/" + key
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
