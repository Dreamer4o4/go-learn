package test

import (
	"golearn/src/httpServer"
	"testing"
)

func findRouter(node *httpServer.TrieTreeRouterNode, paths []string, t *testing.T) {
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
	testnode := httpServer.NewTrieTreeRouterNode("root", nil)

	testnode.InsertNode([]string{"r", "rr"}, func(ctxt *httpServer.Context) {
		t.Log("r-rr")
	})
	testnode.InsertNode([]string{"r", "rr", "rrr"}, func(ctxt *httpServer.Context) {
		t.Log("r-rr-rrr")
	})
	findRouter(testnode, []string{"r", "rr"}, t)
	findRouter(testnode, []string{"r", "rr", "rrr"}, t)

	testnode.InsertNode([]string{"l", "ll", "lll"}, func(ctxt *httpServer.Context) {
		t.Log("l-ll-lll")
	})
	testnode.InsertNode([]string{"l", "ll", "*"}, func(ctxt *httpServer.Context) {
		t.Log("l-ll-*")
	})
	testnode.InsertNode([]string{"l", "lr", "*"}, func(ctxt *httpServer.Context) {
		t.Log("l-lr-*")
	})

	findRouter(testnode, []string{"l", "ll"}, t)
	findRouter(testnode, []string{"l", "ll", "lll"}, t)
	findRouter(testnode, []string{"l", "ll", "t"}, t)
	findRouter(testnode, []string{"l", "lr", "t", "t", "t"}, t)
	findRouter(testnode, []string{"l", "t"}, t)

}
