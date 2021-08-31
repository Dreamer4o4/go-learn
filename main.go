package main

import (
	"fmt"
	"golearn/src/cache"
	cacheStrategy "golearn/src/cache/cachestrategy"
	httpServer "golearn/src/httpserver"
	"golearn/src/util"
)

func HttpServerExample() {
	sh := httpServer.NewHttpServer(":4000")

	sh.AddGetGroupFunc("/", func(ctxt *httpServer.Context) {
		fmt.Println("group fun before: /")
		ctxt.NextStep()
		fmt.Println("group fun after: /")
	})
	sh.StaticServer("/path", "./webpage")
	sh.AddGetFunc("/path", func(ctxt *httpServer.Context) {
		names := []string{"geektutu"}
		fmt.Println(names[100])
	})

	sh.Run()
}

func CacheExample() {
	//	cache server
	db := map[string]string{
		"1": "111",
		"2": "222",
		"3": "333",
	}
	cp := cache.NewCachePool(":4000", nil)
	cp.AddRmoteCache("127.0.0.1:4000", 10)
	cp.AddRmoteCache("127.0.0.1:4001", 10)
	cache.NewCacheMangement(cacheStrategy.NewLruStrategy(100), cp, cache.LoadFunc(func(key string) ([]byte, bool) {
		if v, ok := db[key]; ok {
			fmt.Println("load : ", key)
			return []byte(v), ok
		}
		fmt.Println("query nil : ", key)
		return nil, false
	}))

	//	cache client
	client := cache.NewCachePool("", nil)
	// client.AddRmoteCache("192.168.200.128:4000", 1)
	client.AddRmoteCache("127.0.0.1:4000", 1)
	fmt.Println(client.GetValue("1"))
	fmt.Println(client.GetValue("2"))

}

func Foo() {
	var data struct {
		Name string
		Age  string
	}
	util.ReadJsonConfig("conf/test.json", &data)
	fmt.Println(data)
}

func main() {
	// go HttpServerExample()
	// CacheExample()
	Foo()
}
