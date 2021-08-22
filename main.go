package main

import (
	"fmt"
	"golearn/src/cache"
	cacheStrategy "golearn/src/cache/cachestrategy"
	httpServer "golearn/src/httpserver"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpServerExample() {
	sh := httpServer.NewHttpServer(":4000")

	sh.AddGetGroupFunc("/", func(ctxt *httpServer.Context) {
		fmt.Fprintln(ctxt.Resw, "group fun before: /")
		ctxt.NextStep()
		fmt.Fprintln(ctxt.Resw, "group fun after: /")
	})
	sh.StaticServer("/path", "./webpage")
	sh.AddGetFunc("/path", func(ctxt *httpServer.Context) {
		names := []string{"geektutu"}
		fmt.Println(names[100])
	})

	sh.Run()
}

func CacheExample() {
	lru := cacheStrategy.NewLruStrategy(100)
	cache.NewCacheMangement(lru, cache.NewCachePool(nil), nil)
}

func Foo() {
	time.Sleep(2 * time.Second)
	fmt.Print("--------------\r\n")
	resp, err := http.Get("http://localhost:4000")
	if err != nil {
		fmt.Print("error!!!", err)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("response : %s", bytes)
}

func main() {
	go HttpServerExample()
	// CacheExample()
	Foo()
}
