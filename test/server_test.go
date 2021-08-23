package test

import (
	"golearn/src/httpServer"
	"testing"
)

func TestParasePath(t *testing.T) {
	t.Log(httpServer.ParasePath("get", "/"))
	t.Log(httpServer.ParasePath("get", "/asd/"))
	t.Log(httpServer.ParasePath("get", "/asd"))
	t.Log(httpServer.ParasePath("get", "asd/"))
	t.Log(httpServer.ParasePath("get", "/asd/asd"))
}
