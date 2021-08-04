package src

import (
	"fmt"
	"log"
	"net/http"
)

type handlerfunc func(ctxt *Context)

type serverHandler struct {
	addr  string
	funcs map[string]handlerfunc
}

func NewServerHandler(addr string) *serverHandler {
	return &serverHandler{
		addr:  ":" + addr,
		funcs: make(map[string]handlerfunc),
	}
}

func (sh *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + ":" + r.URL.Path
	if handler, ok := sh.funcs[key]; ok {
		handler(newContext(w, r))
	} else {
		fmt.Fprintln(w, "404 NOT FOUND!\npath : ", r.URL.Path)
	}
}

func (sh *serverHandler) addHandlerFunc(method, path string, handler handlerfunc) {
	key := method + ":" + path
	sh.funcs[key] = handler
}

func (sh *serverHandler) AddGetFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("GET", path, handler)
}

func (sh *serverHandler) AddPostFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("POST", path, handler)
}

func (sh *serverHandler) Run() {
	log.Fatal(http.ListenAndServe(sh.addr, sh))
}
