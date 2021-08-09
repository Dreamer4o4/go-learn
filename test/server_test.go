package test

import (
	"golearn/src/webserver"
	"testing"
)

func TestParasePath(t *testing.T) {
	t.Log(webserver.ParasePath("get", "/"))
	t.Log(webserver.ParasePath("get", "/asd/"))
	t.Log(webserver.ParasePath("get", "/asd"))
	t.Log(webserver.ParasePath("get", "asd/"))
	t.Log(webserver.ParasePath("get", "/asd/asd"))
}
