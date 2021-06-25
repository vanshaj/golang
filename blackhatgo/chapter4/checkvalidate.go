package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type credentials struct {
	username string
	password string
}

func (c *credentials) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) { // create a credentials structure that implements negroni handler interface the http.HandlerFunc represents the next handler func in the middleware chain
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	if username != c.username || password != c.password {
		http.Error(w, "Unauthorized", 401)
		return
	}
	cont := context.WithValue(r.Context(), "username", username)
	r = r.WithContext(cont)
	next(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	fmt.Fprintf(w, "hello dear %s", username)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello).Methods("GET")
	n := negroni.Classic()
	n.Use(&credentials{ // telling negroni to use our custom http handler interface as the middleware chain but that custom handler must implement negroni handler interface
		username: "",
		password: "",
	})
	n.UseHandler(r) // tell negroni to use the router
	http.ListenAndServe(":8000", n)
}
