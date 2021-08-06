package test

import "testing"
import "web/src"

func TestParasePath(t *testing.T) {
	t.Log(src.ParasePath("get", "/"))
	t.Log(src.ParasePath("get", "/asd/"))
	t.Log(src.ParasePath("get", "/asd"))
	t.Log(src.ParasePath("get", "asd/"))
	t.Log(src.ParasePath("get", "/asd/asd"))
}
