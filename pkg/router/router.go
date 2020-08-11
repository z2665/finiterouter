package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/z2665/finiterouter/pkg/tree"
)

type Router struct {
	gettree    *tree.Tree
	posttree   *tree.Tree
	puttree    *tree.Tree
	deletetree *tree.Tree
	isfinished bool
}

func NewRouter() *Router {
	return &Router{
		gettree:    tree.NewTree(),
		posttree:   tree.NewTree(),
		puttree:    tree.NewTree(),
		deletetree: tree.NewTree(),
	}
}
func (r *Router) chooseTree(method string) *tree.Tree {
	switch method {
	case "GET":
		return r.gettree
	case "POST":
		return r.posttree
	case "PUT":
		return r.puttree
	case "DELETE":
		return r.deletetree
	}
	return nil
}
func (r *Router) Done() {
	r.isfinished = true
}
func (r *Router) checkIfFinished() bool {
	return r.isfinished
}
func (r *Router) GET(path string, handler http.HandlerFunc) {
	if r.checkIfFinished() {
		panic(fmt.Sprintf("can't add router after done"))
	}
	paths := strings.Split(path, "/")
	var pathstring string
	for _, k := range paths {
		//只取/后第一位作为路由关键词
		if len(k) > 1 {
			pathstring += string(k[0])
		}
	}
	if len(pathstring) == 0 {
		r.gettree.First().Handle = handler
		return
	}
	if r.gettree.Search(pathstring) != nil {
		panic(fmt.Errorf("get router is exists"))
	}
	r.gettree.Insert(pathstring).Handle = handler
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)
	tr := r.chooseTree(req.Method)
	if tr == nil {
		w.Write([]byte("unsupported method"))
		return
	}
	var pathstring string
	for i, k := range req.URL.Path {
		//只取/后第一位作为路由关键词
		if k == '/' {
			if i+1 < len(req.URL.Path) {
				pathstring += string(req.URL.Path[i+1])
			}
		}
	}
	if len(pathstring) == 0 {
		if tr.First().Handle != nil {
			tr.First().Handle(w, req)
		} else {
			http.NotFound(w, req)
		}
		return
	}
	node := tr.Search(pathstring)
	if node != nil {
		node.Handle(w, req)
	} else {
		http.NotFound(w, req)
	}

}
