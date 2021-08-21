package cache

import (
	"fmt"
	"sort"
)

type Hash func(data []byte) uint32

type ConsistentHash struct {
	hashFunc        Hash
	virtualToReal   map[uint32]*RealServerNode
	realServerNodes map[string]*RealServerNode
	virtualNodesId  []uint32 // circle hash
}

type RealServerNode struct {
	Name           string
	Addr           string
	VirtualNodeNum int
	VirtualNodesId []uint32
}

func NewConsistentHash(hash Hash) *ConsistentHash {
	return &ConsistentHash{
		hashFunc:        hash,
		virtualToReal:   make(map[uint32]*RealServerNode),
		realServerNodes: make(map[string]*RealServerNode),
	}
}

func NewReadServerNode(addr, name string, nodeNum int) *RealServerNode {
	if name == "" {
		name = addr
	}
	if nodeNum <= 0 {
		nodeNum = 3
	}
	return &RealServerNode{
		Name:           name,
		Addr:           addr,
		VirtualNodeNum: nodeNum,
	}
}

func (ch *ConsistentHash) AddRealServer(rs RealServerNode) {
	for idx := 0; idx < rs.VirtualNodeNum; idx++ {
		virtualNodeName := rs.Name + fmt.Sprint(idx)
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
	ch.realServerNodes[rs.Name] = &rs

	sort.Slice(ch.virtualNodesId, func(i, j int) bool {
		return ch.virtualNodesId[i] < ch.virtualNodesId[j]
	})
}

/*
**	delete realserver
**	ConsistentHash only delete realServerNodes & virtualToReal, virtualNodesId lazy delete
**	only delete realServer now, not move cache data
 */
func (ch *ConsistentHash) RemoveRealServer(serverName string) {
	if realServer, ok := ch.realServerNodes[serverName]; ok {
		delete(ch.realServerNodes, serverName)
		for _, virtualNodeId := range realServer.VirtualNodesId {
			delete(ch.virtualToReal, virtualNodeId)
		}
	}
}

/*
**	find RealServer has k-v data
**	lazy delete useless virtualNodesId
 */
func (ch *ConsistentHash) FindServer(key string) *RealServerNode {
	if len(ch.virtualNodesId) == 0 {
		return nil
	}

	keyHash := ch.hashFunc([]byte(key))
	for {
		idx := sort.Search(len(ch.virtualNodesId), func(i int) bool {
			return ch.virtualNodesId[i] >= keyHash
		})
		idx = idx % len(ch.virtualNodesId)
		if realSerevr, ok := ch.virtualToReal[ch.virtualNodesId[idx]]; ok {
			return realSerevr
		} else {
			ch.virtualNodesId = append(ch.virtualNodesId[:idx], ch.virtualNodesId[idx+1:]...)
		}

	}

}
