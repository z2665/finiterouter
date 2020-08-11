package main

import (
	"fmt"
	"net/http"

	router "github.com/z2665/finiterouter/pkg/router"
)

func main() {

	ro := router.NewRouter()
	ro.GET("/hello/mark", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("GET %s\n", req.RequestURI)))
	})
	ro.GET("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(fmt.Sprintf("index\n")))
	})

	http.ListenAndServe(":8080", ro)
}
