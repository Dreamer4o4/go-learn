package test

import (
	"golearn/src/util"
	"testing"
)

type TestConf struct {
	Name string
	Age  string
}

func TestJsonConfig(t *testing.T) {
	var data TestConf
	util.ReadJsonConfig("../conf/test.json", &data)
	t.Log(data)
}
