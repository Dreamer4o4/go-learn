package test

import (
	"testing"
	"web/src"
)

func TestTrieTreeNode(t *testing.T) {
	node := src.NewTrieTreeNode("hello")
	t.Log(node.Value())
}
