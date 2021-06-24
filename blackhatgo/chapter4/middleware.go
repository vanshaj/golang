package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hey man how are you")
}

type router struct{}

func (rou *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/a":
		fmt.Fprintf(w, "you have been entered to a")
	case "/b":
		fmt.Fprintf(w, "you have been entered to b")
	default:
		http.Error(w, "Not found", 404)
	}
}

type logger struct {
	Inner http.Handler
}

func (l *logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("start")
	l.Inner.ServeHTTP(w, req)
	log.Println("end")
}

func main() {
	//http.HandleFunc("/hello", hello)

	//var r router
	//http.ListenAndServe(":8000", &r)

	h := http.HandlerFunc(hello)
	l := logger{Inner: h}
	http.ListenAndServe(":8000", &l)
}
