package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"    // router
	"github.com/urfave/negroni" // middleware to perform auth on every request
)

func main() {
	r := mux.NewRouter()
	n := negroni.Classic() // creates pointer to a negroni instance
	n.UseHandler(r)
	// r.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
	// 	io.WriteString(w, "you have reached bye")
	// }).Methods("GET")

	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		user := mux.Vars(r)["name"]
		fmt.Fprintf(w, "hello %s", user)
	}).Methods("GET")
	http.ListenAndServe(":8000", n) //r.HandleFunc("/users/{user:[a-z]+}", can use regex also to mention the case sensitivity
}
