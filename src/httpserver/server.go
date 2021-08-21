package httpServer

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type handlerfunc func(ctxt *Context)

const rootPath string = "root"

type httpServer struct {
	ipAddr   string
	rootPath *TrieTreeRouterNode
}

func NewHttpServer(ipAddr string) *httpServer {
	return &httpServer{
		ipAddr:   ipAddr,
		rootPath: NewTrieTreeRouterNode(rootPath, nil),
	}
}

func ParasePath(method, path string) []string {
	var retPath []string

	retPath = append(retPath, method)

	// remove '/' at two side of path
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

func (sh *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths := ParasePath(r.Method, r.URL.Path)
	if handler, groupHandlers := sh.rootPath.FindHanlder(paths); handler != nil {
		context := newContext(w, r)

		// load request handle functions
		context.AddSteps(groupHandlers...)
		context.AddSteps(handler)

		// start
		context.NextStep()
	} else {
		// wrong request
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 NOT FOUND!\npath : ", r.URL.Path)
	}
}

func (sh *httpServer) addHandlerFunc(method, path string, handler handlerfunc) {
	sh.rootPath.InsertHandler(ParasePath(method, path), handler)
}

func (sh *httpServer) addGroupHandlerFunc(method, path string, groupHandler handlerfunc) {
	sh.rootPath.InsertGroupHandlers(ParasePath(method, path), groupHandler)
}

func (sh *httpServer) AddGetFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("GET", path, handler)
}

func (sh *httpServer) AddPostFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("POST", path, handler)
}

func (sh *httpServer) AddGetGroupFunc(path string, groupHandler handlerfunc) {
	sh.addGroupHandlerFunc("GET", path, groupHandler)
}

func (sh *httpServer) AddPostGroupFunc(path string, groupHandler handlerfunc) {
	sh.addGroupHandlerFunc("POST", path, groupHandler)
}

func (sh *httpServer) StaticServer(urlPath, srcPath string) {
	sh.AddGetFunc(urlPath+"/*", func(ctxt *Context) {
		http.StripPrefix(urlPath, http.FileServer(http.Dir(srcPath))).ServeHTTP(ctxt.Resw, ctxt.Req)
	})
}

func (sh *httpServer) GobalHandler(handlers ...handlerfunc) {
	sh.rootPath.groupHandler = append(sh.rootPath.groupHandler, handlers...)
}

func (sh *httpServer) Run() {
	sh.GobalHandler(logger(), recoverHandler())
	// sh.rootPath.show()	//for debug
	log.Fatal(http.ListenAndServe(sh.ipAddr, sh))
}
