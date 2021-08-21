package cache

type CacheMangement struct {
	hashMap  ConsistentHash
	strategy CacheStrategy
}
