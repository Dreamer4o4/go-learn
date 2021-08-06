package src

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type handlerfunc func(ctxt *Context)

type serverHandler struct {
	ipaddr   string
	rootPath *TrieTreeRouterNode
}

func NewServerHandler(ipaddr string) *serverHandler {
	return &serverHandler{
		ipaddr:   ipaddr,
		rootPath: NewTrieTreeRouterNode("root", nil),
	}
}

func parasePath(method, path string) []string {
	var retPath []string

	retPath = append(retPath, method)

	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	if len(path) != 0 {
		retPath = append(retPath, strings.Split(path, "/")...)
	}
	return retPath
}

func (sh *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths := parasePath(r.Method, r.URL.Path)
	if handler, groupHandlers := sh.rootPath.FindHanlder(paths); handler != nil {
		context := newContext(w, r)
		context.AddSteps(groupHandlers...)
		context.AddSteps(handler)

		context.NextStep()
	} else {
		fmt.Fprintln(w, "404 NOT FOUND!\npath : ", r.URL.Path)
	}
}

func (sh *serverHandler) addHandlerFunc(method, path string, handler handlerfunc) {
	sh.rootPath.InsertHandler(parasePath(method, path), handler)
}

func (sh *serverHandler) addGroupHandlerFunc(method, path string, groupHandler handlerfunc) {
	sh.rootPath.InsertGroupHandlers(parasePath(method, path), groupHandler)
}

func (sh *serverHandler) AddGetFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("GET", path, handler)
}

func (sh *serverHandler) AddPostFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("POST", path, handler)
}

func (sh *serverHandler) AddGetGroupFunc(path string, groupHandler handlerfunc) {
	sh.addGroupHandlerFunc("GET", path, groupHandler)
}

func (sh *serverHandler) AddPostGroupFunc(path string, groupHandler handlerfunc) {
	sh.addGroupHandlerFunc("POST", path, groupHandler)
}

func (sh *serverHandler) Run() {
	sh.rootPath.show()
	log.Fatal(http.ListenAndServe(sh.ipaddr, sh))
}
