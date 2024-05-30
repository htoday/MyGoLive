package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8081", "http service address")
var house = make(map[string]*Hub)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "app/talkService/home.html")
}
func main() {
	flag.Parse()
	hub := NewHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws/{room}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		room := vars["room"]
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
