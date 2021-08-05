package src

type TrieTreeRouterNode struct {
	path      string
	handler   handlerfunc
	nextNodes map[string]*TrieTreeRouterNode
}

func NewTrieTreeRouterNode(path string, handler handlerfunc) *TrieTreeRouterNode {
	return &TrieTreeRouterNode{
		path:      path,
		handler:   handler,
		nextNodes: make(map[string]*TrieTreeRouterNode),
	}
}

func (node *TrieTreeRouterNode) insertNextOneNode(path string, handler handlerfunc) *TrieTreeRouterNode {
	if _, exist := node.nextNodes[path]; !exist {
		node.nextNodes[path] = NewTrieTreeRouterNode(path, handler)
	}

	return node.nextNodes[path]
}

func (node *TrieTreeRouterNode) findNextOneNode(path string) *TrieTreeRouterNode {
	if res, exist := node.nextNodes[path]; exist {
		return res
	} else if res, exist := node.nextNodes["*"]; exist {
		return res
	}

	return nil
}

func (node *TrieTreeRouterNode) InsertNode(paths []string, handler handlerfunc) *TrieTreeRouterNode {
	curNode := node
	for index, curPath := range paths {
		if index == len(paths)-1 || curPath == "*" {
			curNode = curNode.insertNextOneNode(curPath, handler)
			break
		}

		curNode = curNode.insertNextOneNode(curPath, nil)
	}
	return curNode
}

func (node *TrieTreeRouterNode) FindNode(paths []string) *TrieTreeRouterNode {
	curNode := node
	for _, curPath := range paths {
		curNode = curNode.findNextOneNode(curPath)
		if curNode == nil || curNode.path == "*" {
			break
		}
	}
	return curNode
}

func (node *TrieTreeRouterNode) Handler() handlerfunc {
	return node.handler
}
