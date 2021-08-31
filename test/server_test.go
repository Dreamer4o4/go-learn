package test

import (
	"golearn/src/httpServer"
	"testing"
)

func TestParasePathWithMethod(t *testing.T) {
	t.Log(httpServer.ParasePathWithMethod("get", "/"))
	t.Log(httpServer.ParasePathWithMethod("get", "/asd/"))
	t.Log(httpServer.ParasePathWithMethod("get", "/asd"))
	t.Log(httpServer.ParasePathWithMethod("get", "asd/"))
	t.Log(httpServer.ParasePathWithMethod("get", "/asd/asd"))
}
