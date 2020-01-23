package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	h "./handlers"
)

func main() {
	rout := mux.NewRouter()
	rout.HandleFunc("/queue", h.QueueHandler)

	log.Fatal(http.ListenAndServe(":80", rout))
}
