package main

import cacheStrategy "golearn/src/cache/cacheStrategy"

func main() {
	// sh := webserver.NewWebServer(":4000")

	// sh.AddGetGroupFunc("/", func(ctxt *webserver.Context) {
	// 	fmt.Fprintln(ctxt.Resw, "group fun before: /")
	// 	ctxt.NextStep()
	// 	fmt.Fprintln(ctxt.Resw, "group fun after: /")
	// })
	// sh.StaticServer("/path", "./webpage")
	// sh.AddGetFunc("/path", func(ctxt *webserver.Context) {
	// 	names := []string{"geektutu"}
	// 	fmt.Println(names[100])
	// })

	// sh.Run()

	test := cacheStrategy.NewLruStrategy(100)
	str := "asd"
	test.Push("t1", cacheStrategy.NewByteValue(str))
	test.Push("t1", cacheStrategy.NewByteValue(str))
	test.Push("t", cacheStrategy.NewByteValue(str))

}
