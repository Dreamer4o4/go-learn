package main

import (
	"fmt"
	"golearn/src/httpServer"
)

func main() {
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

	// test := cacheStrategy.NewLruStrategy(10)
	// str := "asd"
	// test.Push("t1", cacheStrategy.NewByteValue(str))
	// test.Push("t2", cacheStrategy.NewByteValue(str))
	// test.Push("t", cacheStrategy.NewByteValue(str))

	// if _, ok := test.Find("t"); ok {
	// 	fmt.Print("find t")
	// }

	// fmt.Print(test.Pop("t"))
	// fmt.Print(test.Pop("t"))

	// if _, ok := test.Find("t"); ok {
	// 	fmt.Print("find t")
	// }
	// if _, ok := test.Find("t1"); ok {
	// 	fmt.Print("find t1")
	// }

}
