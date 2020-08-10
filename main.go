package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"

	"github.com/gorilla/mux"
)

var indexTmpl *template.Template

var db []byte

func init() {
	indexTmpl = template.Must(template.ParseFiles(
		path.Join("views", "index.html"),
	))

	// Read shedule.json db
	file, _ := os.Open("./schedule.json")
	db, _ = ioutil.ReadAll(file)

}

func main() {
	r := mux.NewRouter()

	// Index page
	r.HandleFunc("/", indexHandler).Methods("GET").Schemes("http")

	// API
	r.HandleFunc("/api/time/", apiTimeHandler).Methods("GET").Schemes("http")
	r.HandleFunc("/api/db/", apiDbHandler).Methods("GET").Schemes("http")
	r.HandleFunc("/api/lessons/", apiLessonsHandler).Queries(
		"group", "{group:(?:1|2)}",
		"day", "{day:(?:yesterday|today|tomorrow)}",
	).Methods("GET").Schemes("http")

	// Setup static folder
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))),
	)

	http.Handle("/", r)

	// Run
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
