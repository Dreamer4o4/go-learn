package test

import (
	"testing"
	"web/src"
)

func findRouter(node *src.TrieTreeRouterNode, paths []string, t *testing.T) {
	funVar := node.FindNode(paths)
	if funVar == nil {
		t.Log("nil node", paths)
	} else if funVar.Handler() == nil {
		t.Log("nil func", paths)
	} else {
		funVar.Handler()(nil)
	}
}

func TestTrieTreeNode(t *testing.T) {
	testnode := src.NewTrieTreeRouterNode("root", nil)
	testnode.InsertNode([]string{"l", "ll", "lll"}, func(ctxt *src.Context) {
		t.Log("l-ll-lll")
	})
	testnode.InsertNode([]string{"l", "ll", "*"}, func(ctxt *src.Context) {
		t.Log("l-ll-*")
	})
	testnode.InsertNode([]string{"l", "lr", "*"}, func(ctxt *src.Context) {
		t.Log("l-lr-*")
	})

	findRouter(testnode, []string{"l", "ll"}, t)
	findRouter(testnode, []string{"l", "ll", "lll"}, t)
	findRouter(testnode, []string{"l", "ll", "t"}, t)
	findRouter(testnode, []string{"l", "lr", "t", "t", "t"}, t)
	findRouter(testnode, []string{"l", "t"}, t)

}
