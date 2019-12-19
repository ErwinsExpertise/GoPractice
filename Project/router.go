package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(port string) {

	rout := mux.NewRouter()

	rout.HandleFunc("/", HomeHandler)
	rout.HandleFunc("/shark", SharkHandler)

	log.Fatal(http.ListenAndServe(port, rout))
}
