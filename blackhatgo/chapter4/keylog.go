package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var listenAddr string
var wsAddr string
var jsTemplate *template.Template
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	flag.StringVar(&listenAddr, "listen address", "", "Adress")
	flag.StringVar(&wsAddr, "websocket", "", "ws")
	flag.Parse()
	var err error
	jsTemplate, err = template.ParseFiles("logges.js")
	if err != nil {
		panic(err)
	}
}

func serveJS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "", 500)
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Printf("From %s: %s\n", conn.RemoteAddr().String(), string(msg))
	}

}

func serveFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	jsTemplate.Execute(w, wsAddr)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ws", serveJS)
	r.HandleFunc("/k.js", serveFile)
	log.Fatal(http.ListenAndServe(":9090", r))
}
