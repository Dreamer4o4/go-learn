package test

import (
	"golearn/src/cache"
	"strconv"
	"testing"
)

func TestConsistentHash(t *testing.T) {
	hashMap := cache.NewConsistentHash(func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
		"57": "6",
		"41": "8",
	}
	hashMap.AddRealServer(*cache.NewRealServerNode("6", 1))
	hashMap.AddRealServer(*cache.NewRealServerNode("4", 2))
	hashMap.AddRealServer(*cache.NewRealServerNode("2", 1))

	for key, _ := range testCases {
		t.Log("cur key : ", key)
		t.Log("find server : ", hashMap.FindServer(key))
	}
}
