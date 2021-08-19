package cache

import (
	"fmt"
	"sort"
)

type Hash func(data []byte) uint32

type ConsistentHash struct {
	hashFunc       Hash
	virtualToReal  map[uint32]*RealSeverNode
	virtualNodesId []uint32
}

type RealSeverNode struct {
	Name           string
	Addr           string
	VirtualNodeNum int
	VirtualNodesId []uint32
}

func NewConsistentHash(hash Hash) *ConsistentHash {
	return &ConsistentHash{
		hashFunc:      hash,
		virtualToReal: make(map[uint32]*RealSeverNode),
	}
}

func NewReadServerNode(addr, name string, nodeNum int) *RealSeverNode {
	return &RealSeverNode{
		Name:           name,
		Addr:           addr,
		VirtualNodeNum: nodeNum,
	}
}

func (ch *ConsistentHash) AddReadServer(rs RealSeverNode) {
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

	sort.Slice(ch.virtualNodesId, func(i, j int) bool {
		return ch.virtualNodesId[i] < ch.virtualNodesId[j]
	})
}
