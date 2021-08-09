package test

import (
	"golearn/src/webserver"
	"testing"
)

func findRouter(node *webserver.TrieTreeRouterNode, paths []string, t *testing.T) {
	funVar, _ := node.FindNode(paths)
	if funVar == nil {
		t.Log("nil node", paths)
	} else if funVar.Handler() == nil {
		t.Log("nil func", paths)
	} else {
		funVar.Handler()(nil)
	}
}

func TestTrieTreeNode(t *testing.T) {
	testnode := webserver.NewTrieTreeRouterNode("root", nil)
	testnode.InsertNode([]string{"l", "ll", "lll"}, func(ctxt *webserver.Context) {
		t.Log("l-ll-lll")
	})
	testnode.InsertNode([]string{"l", "ll", "*"}, func(ctxt *webserver.Context) {
		t.Log("l-ll-*")
	})
	testnode.InsertNode([]string{"l", "lr", "*"}, func(ctxt *webserver.Context) {
		t.Log("l-lr-*")
	})

	findRouter(testnode, []string{"l", "ll"}, t)
	findRouter(testnode, []string{"l", "ll", "lll"}, t)
	findRouter(testnode, []string{"l", "ll", "t"}, t)
	findRouter(testnode, []string{"l", "lr", "t", "t", "t"}, t)
	findRouter(testnode, []string{"l", "t"}, t)

}
