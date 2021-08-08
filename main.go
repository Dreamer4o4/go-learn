package main

import (
	"fmt"
	"web/src"
)

func main() {
	sh := src.NewServerHandler(":4000")
	// sh.AddGetFunc("/patha/*", func(ctxt *src.Context) {
	// 	ctxt.ResHtml("hello.html")
	// 	ctxt.NextStep()
	// 	fmt.Fprintln(ctxt.Resw, "normal fun : /patha/*")
	// })
	// sh.AddGetFunc("/pathb", func(ctxt *src.Context) {
	// 	fmt.Fprintln(ctxt.Resw, "normal fun : /pathb")
	// })
	// sh.AddGetGroupFunc("/", func(ctxt *src.Context) {
	// 	fmt.Fprintln(ctxt.Resw, "group fun before: /")
	// 	ctxt.NextStep()
	// 	fmt.Fprintln(ctxt.Resw, "group fun after: /")
	// })
	// sh.AddGetGroupFunc("/patha/sec/", func(ctxt *src.Context) {
	// 	fmt.Fprintln(ctxt.Resw, "group fun before: /patha/sec/")
	// 	ctxt.NextStep()
	// 	fmt.Fprintln(ctxt.Resw, "group fun after: /patha/sec/")
	// })

	sh.StaticServer("/path", "./webpage")

	sh.AddGetFunc("/path", func(ctxt *src.Context) {
		names := []string{"geektutu"}
		fmt.Println(names[100])
	})
	sh.Run()
}
