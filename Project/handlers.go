package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Hostname string
}

type Fact struct {
	Index   int
	Content string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var page Page
	hostname, err := os.Hostname()

	if err != nil {
		log.Println(err)
	}

	page.Hostname = hostname

	tmpl := template.Must(template.ParseFiles("home.html"))

	tmpl.Execute(w, page)
}

func SharkHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("shark.html"))

	tmpl.Execute(w, nil)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	var facts []Fact

	host := ""
	port := ""
	user := ""
	password := ""
	schema := ""
	db, err := sql.Open("mysql", user+":"+password+"@"+host+":"+port+"/"+schema)

	if err != nil {
		log.Println(err)
	}
	result := db.QueryRow("SELECT * FROM facts").Scan()

	if err != nil {
		log.Println(err)
	}

	log.Println(result)
	log.Println(facts)

}
