package main

import (
	"web/src"
)

func main() {
	sh := src.NewServerHandler("4000")
	sh.AddGetFunc("/", func(ctxt *src.Context) {
		ctxt.ResHtml("hello.html")
	})
	sh.Run()
}
