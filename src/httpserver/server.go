package httpServer

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type handlerfunc func(ctxt *Context)

const rootPath string = "root"

type HttpServer struct {
	ipAddr   string
	rootPath *TrieTreeRouterNode
}

func NewHttpServer(ipAddr string) *HttpServer {
	return &HttpServer{
		ipAddr:   ipAddr,
		rootPath: NewTrieTreeRouterNode(rootPath, nil),
	}
}

func ParasePath(path string) []string {
	var retPath []string

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

func ParasePathWithMethod(method, path string) []string {
	var retPath []string

	retPath = append(retPath, method)
	retPath = append(retPath, ParasePath(path)...)

	return retPath
}

func (sh *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths := ParasePathWithMethod(r.Method, r.URL.Path)
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

func (sh *HttpServer) addHandlerFunc(method, path string, handler handlerfunc) {
	sh.rootPath.InsertHandler(ParasePathWithMethod(method, path), handler)
}

func (sh *HttpServer) addGroupHandlerFunc(method, path string, groupHandler handlerfunc) {
	sh.rootPath.InsertGroupHandlers(ParasePathWithMethod(method, path), groupHandler)
}

func (sh *HttpServer) AddGetFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("GET", path, handler)
}

func (sh *HttpServer) AddPostFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("POST", path, handler)
}

func (sh *HttpServer) AddGetGroupFunc(path string, groupHandler handlerfunc) {
	sh.addGroupHandlerFunc("GET", path, groupHandler)
}

func (sh *HttpServer) AddPostGroupFunc(path string, groupHandler handlerfunc) {
	sh.addGroupHandlerFunc("POST", path, groupHandler)
}

func (sh *HttpServer) StaticServer(urlPath, srcPath string) {
	sh.AddGetFunc(urlPath+"/*", func(ctxt *Context) {
		http.StripPrefix(urlPath, http.FileServer(http.Dir(srcPath))).ServeHTTP(ctxt.Resw, ctxt.Req)
	})
}

func (sh *HttpServer) GobalHandler(handlers ...handlerfunc) {
	sh.rootPath.groupHandler = append(sh.rootPath.groupHandler, handlers...)
}

func (sh *HttpServer) Run() {
	sh.GobalHandler(logger(), recoverHandler())
	// sh.rootPath.show() //for debug
	log.Fatal(http.ListenAndServe(sh.ipAddr, sh))
}
