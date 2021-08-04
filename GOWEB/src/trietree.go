package src

type TrieTreeNode struct {
	value     string
	nextNodes map[string]*TrieTreeNode
}

func NewTrieTreeNode(value string) *TrieTreeNode {
	return &TrieTreeNode{
		value:     value,
		nextNodes: make(map[string]*TrieTreeNode),
	}
}

func (node *TrieTreeNode) InsertOneNode(value string) *TrieTreeNode {
	if _, exist := node.nextNodes[value]; !exist {
		node.nextNodes[value] = NewTrieTreeNode(value)
	}

	return node.nextNodes[value]
}

func (node *TrieTreeNode) FindNextOneNode(value string) *TrieTreeNode {
	if res, exist := node.nextNodes[value]; exist {
		return res
	}

	return nil
}

func (node *TrieTreeNode) InsertNodes(values []string) *TrieTreeNode {
	curNode := node
	for _, value := range values {
		curNode.InsertOneNode(value)
		curNode = curNode.FindNextOneNode(value)
		if curNode == nil {
			panic("trie tree insert error")
		}
	}
	return curNode
}

func (node *TrieTreeNode) FindNode(values []string) *TrieTreeNode {
	curNode := node
	for _, value := range values {
		curNode = curNode.FindNextOneNode(value)
		if curNode == nil {
			return nil
		}
	}
	return curNode
}

func (node *TrieTreeNode) Value() string {
	return node.value
}
