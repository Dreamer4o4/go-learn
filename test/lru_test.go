package test

import (
	"golearn/src/cache/cacheStrategy"
	"testing"
)

func TestLruStrategy(t *testing.T) {
	test := cacheStrategy.NewLruStrategy(10)
	str := "asd"
	test.Push("t1", cacheStrategy.NewByteValue(str))
	test.Push("t2", cacheStrategy.NewByteValue(str))
	test.Push("t", cacheStrategy.NewByteValue(str))

	if _, ok := test.Find("t"); ok {
		t.Log("find t")
	}

	t.Log(test.Pop("t"))
	t.Log(test.Pop("t"))

	if _, ok := test.Find("t"); ok {
		t.Log("find t")
	}
	if _, ok := test.Find("t1"); ok {
		t.Log("find t1")
	}
}
