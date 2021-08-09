package main

import (
	"fmt"
	"golearn/src/webserver"
)

func main() {
	sh := webserver.NewWebServer(":4000")
	// sh.AddGetFunc("/patha/*", func(ctxt *webserver.Context) {
	// 	ctxt.ResHtml("hello.html")
	// 	ctxt.NextStep()
	// 	fmt.Fprintln(ctxt.Resw, "normal fun : /patha/*")
	// })
	// sh.AddGetFunc("/pathb", func(ctxt *webserver.Context) {
	// 	fmt.Fprintln(ctxt.Resw, "normal fun : /pathb")
	// })
	// sh.AddGetGroupFunc("/", func(ctxt *webserver.Context) {
	// 	fmt.Fprintln(ctxt.Resw, "group fun before: /")
	// 	ctxt.NextStep()
	// 	fmt.Fprintln(ctxt.Resw, "group fun after: /")
	// })
	// sh.AddGetGroupFunc("/patha/sec/", func(ctxt *webserver.Context) {
	// 	fmt.Fprintln(ctxt.Resw, "group fun before: /patha/sec/")
	// 	ctxt.NextStep()
	// 	fmt.Fprintln(ctxt.Resw, "group fun after: /patha/sec/")
	// })

	sh.StaticServer("/path", "./webpage")

	sh.AddGetFunc("/path", func(ctxt *webserver.Context) {
		names := []string{"geektutu"}
		fmt.Println(names[100])
	})
	sh.Run()
}
