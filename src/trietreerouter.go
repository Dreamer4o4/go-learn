package src

import "fmt"

type TrieTreeRouterNode struct {
	path         string
	handler      handlerfunc
	groupHandler []handlerfunc
	nextNodes    map[string]*TrieTreeRouterNode
}

func NewTrieTreeRouterNode(path string, handler handlerfunc) *TrieTreeRouterNode {
	return &TrieTreeRouterNode{
		path:         path,
		handler:      handler,
		groupHandler: make([]handlerfunc, 0),
		nextNodes:    make(map[string]*TrieTreeRouterNode),
	}
}

func (node *TrieTreeRouterNode) insertNextNode(path string, handler handlerfunc) *TrieTreeRouterNode {
	if _, exist := node.nextNodes[path]; !exist {
		node.nextNodes[path] = NewTrieTreeRouterNode(path, handler)
	}

	return node.nextNodes[path]
}

func (node *TrieTreeRouterNode) findNextNode(path string) *TrieTreeRouterNode {
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
			curNode = curNode.insertNextNode(curPath, handler)
			break
		}

		curNode = curNode.insertNextNode(curPath, nil)
	}
	return curNode
}

func (node *TrieTreeRouterNode) FindNode(paths []string) (*TrieTreeRouterNode, []handlerfunc) {
	curNode := node
	var groupHandler []handlerfunc

	for _, curPath := range paths {
		if curNode.groupHandler != nil {
			groupHandler = append(groupHandler, curNode.groupHandler...)
		}
		curNode = curNode.findNextNode(curPath)
		if curNode == nil || curNode.path == "*" {
			break
		}
	}
	if curNode != nil && curNode.groupHandler != nil {
		groupHandler = append(groupHandler, curNode.groupHandler...)
	}
	return curNode, groupHandler
}

func (node *TrieTreeRouterNode) Handler() handlerfunc {
	return node.handler
}

func (node *TrieTreeRouterNode) InsertHandler(paths []string, handler handlerfunc) {
	node.InsertNode(paths, handler)
}

func (node *TrieTreeRouterNode) FindHanlder(paths []string) (handlerfunc, []handlerfunc) {
	targetNode, groupHandler := node.FindNode(paths)
	if targetNode != nil && targetNode.Handler() != nil {
		return targetNode.Handler(), groupHandler
	}
	return nil, nil
}

func (node *TrieTreeRouterNode) InsertGroupHandlers(paths []string, groupHandler handlerfunc) {
	targetNode, _ := node.FindNode(paths)
	if targetNode != nil && targetNode.groupHandler == nil {
		targetNode.groupHandler = append(targetNode.groupHandler, groupHandler)
	}
}

func (node *TrieTreeRouterNode) show() {
	curNodes := []*TrieTreeRouterNode{node}
	curNum := 1
	for curNum != 0 {
		var nextNodes []*TrieTreeRouterNode
		nextNum := 0
		for _, curNode := range curNodes {
			if curNode != nil {
				fmt.Print(curNode.path)
				if curNode.handler != nil {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
				if curNode.groupHandler != nil {
					fmt.Print("1")
				} else {
					fmt.Print("0")
				}
				fmt.Print(" ")
			} else {
				fmt.Print("NULL" + " ")
			}

			for _, v := range curNode.nextNodes {
				nextNodes = append(nextNodes, v)
				if v != nil {
					nextNum++
				}
			}
		}
		fmt.Printf("\r\n")

		curNodes = nextNodes
		curNum = nextNum
	}

}
