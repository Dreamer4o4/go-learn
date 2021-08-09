package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type handlerfunc func(ctxt *Context)

const rootPath string = "root"

type webServer struct {
	ipAddr   string
	rootPath *TrieTreeRouterNode
}

func NewWebServer(ipAddr string) *webServer {
	return &webServer{
		ipAddr:   ipAddr,
		rootPath: NewTrieTreeRouterNode(rootPath, nil),
	}
}

func ParasePath(method, path string) []string {
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

func (sh *webServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paths := ParasePath(r.Method, r.URL.Path)
	if handler, groupHandlers := sh.rootPath.FindHanlder(paths); handler != nil {
		context := newContext(w, r)
		context.AddSteps(groupHandlers...)
		context.AddSteps(handler)

		context.NextStep()
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "404 NOT FOUND!\npath : ", r.URL.Path)
	}
}

func (sh *webServer) addHandlerFunc(method, path string, handler handlerfunc) {
	sh.rootPath.InsertHandler(ParasePath(method, path), handler)
}

func (sh *webServer) addGroupHandlerFunc(method, path string, groupHandler handlerfunc) {
	sh.rootPath.InsertGroupHandlers(ParasePath(method, path), groupHandler)
}

func (sh *webServer) AddGetFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("GET", path, handler)
}

func (sh *webServer) AddPostFunc(path string, handler handlerfunc) {
	sh.addHandlerFunc("POST", path, handler)
}

func (sh *webServer) AddGetGroupFunc(path string, groupHandler handlerfunc) {
	sh.addGroupHandlerFunc("GET", path, groupHandler)
}

func (sh *webServer) AddPostGroupFunc(path string, groupHandler handlerfunc) {
	sh.addGroupHandlerFunc("POST", path, groupHandler)
}

func (sh *webServer) StaticServer(urlPath, srcPath string) {
	sh.AddGetFunc(urlPath+"/*", func(ctxt *Context) {
		http.StripPrefix(urlPath, http.FileServer(http.Dir(srcPath))).ServeHTTP(ctxt.Resw, ctxt.Req)
	})
}

func (sh *webServer) GobalHandler(handlers ...handlerfunc) {
	sh.rootPath.groupHandler = append(sh.rootPath.groupHandler, handlers...)
}

func (sh *webServer) Run() {
	sh.GobalHandler(logger(), recoverHandler())
	// sh.rootPath.show()	//for debug
	log.Fatal(http.ListenAndServe(sh.ipAddr, sh))
}
