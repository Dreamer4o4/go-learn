package cache

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Hash func(data []byte) uint32

const defultVirtualNodeNum int = 10

type ConsistentHash struct {
	hashFunc        Hash
	virtualToReal   map[uint32]*realServerNode
	realServerNodes map[string]*realServerNode
	virtualNodesId  []uint32 // circle hash
}

type realServerNode struct {
	Addr           string
	VirtualNodeNum int
	VirtualNodesId []uint32
}

func NewConsistentHash(hash Hash) *ConsistentHash {
	if hash == nil {
		hash = crc32.ChecksumIEEE
	}

	return &ConsistentHash{
		hashFunc:        hash,
		virtualToReal:   make(map[uint32]*realServerNode),
		realServerNodes: make(map[string]*realServerNode),
	}
}

func NewRealServerNode(addr string, nodeNum int) *realServerNode {
	if nodeNum <= 0 {
		nodeNum = defultVirtualNodeNum
	}
	return &realServerNode{
		Addr:           addr,
		VirtualNodeNum: nodeNum,
	}
}

func (ch *ConsistentHash) AddRealServer(rs realServerNode) {
	for idx := 0; idx < rs.VirtualNodeNum; idx++ {
		virtualNodeName := rs.Addr + fmt.Sprint(idx)
		virtualNodeId := ch.hashFunc([]byte(virtualNodeName))

		for {
			_, ok := ch.virtualToReal[virtualNodeId]
			if !ok {
				break
			}
			virtualNodeId += 10
		}

		ch.virtualToReal[virtualNodeId] = &rs
		ch.virtualNodesId = append(ch.virtualNodesId, virtualNodeId)
		rs.VirtualNodesId = append(rs.VirtualNodesId, virtualNodeId)
	}
	ch.realServerNodes[rs.Addr] = &rs

	sort.Slice(ch.virtualNodesId, func(i, j int) bool {
		return ch.virtualNodesId[i] < ch.virtualNodesId[j]
	})
}

/*
**	delete realserver
**	ConsistentHash only delete realServerNodes & virtualToReal, virtualNodesId lazy delete
**	only delete realServer now, not move cache data
 */
func (ch *ConsistentHash) RemoveRealServer(serverAddr string) {
	if realServer, ok := ch.realServerNodes[serverAddr]; ok {
		delete(ch.realServerNodes, serverAddr)
		for _, virtualNodeId := range realServer.VirtualNodesId {
			delete(ch.virtualToReal, virtualNodeId)
		}
	}
}

/*
**	find RealServer has k-v data
**	lazy delete useless virtualNodesId
 */
func (ch *ConsistentHash) FindServer(key string) string {
	if len(ch.virtualNodesId) == 0 {
		return ""
	}

	keyHash := ch.hashFunc([]byte(key))
	for {
		idx := sort.Search(len(ch.virtualNodesId), func(i int) bool {
			return ch.virtualNodesId[i] >= keyHash
		})
		idx = idx % len(ch.virtualNodesId)
		if realServer, ok := ch.virtualToReal[ch.virtualNodesId[idx]]; ok {
			return realServer.Addr
		} else {
			ch.virtualNodesId = append(ch.virtualNodesId[:idx], ch.virtualNodesId[idx+1:]...)
		}

	}

}
