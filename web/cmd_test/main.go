package main

import (
	"fmt"
	"github.com/herrhu97/simple-go-framework/log"
	"net/http"
)

func main() {
	if false {
		http.HandleFunc("/", indexHandler)
		http.HandleFunc("/hello", helloHandler)
		log.Fatal(http.ListenAndServe(":9999", nil))
	}
	if true {
		engine := new(Engine)
		log.Fatal(http.ListenAndServe(":9999", engine))
	}

}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

// Engine is the uni handler for all requests
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
