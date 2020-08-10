package main

import (
	"fmt"
	"net/http"

	"github.com/tidwall/gjson"
)

// Index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

// API for get times of lessons
func apiTimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	jsonResp := gjson.Get(string(db), "time")
	fmt.Fprintf(w, "%s", jsonResp.String())
}

// API for get full DB of lessons
func apiDbHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", string(db))
}

// API for get lessons for group (1 | 2) and day (yesterday | today | tomorrow)
func apiLessonsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	group := params.Get("group")
	day := params.Get("day")

	star, weekday := getStarAndDay(day)

	gjsonString := fmt.Sprintf("group%s.%s.%s", group, weekday, star)
	jsonResp := gjson.Get(string(db), gjsonString)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", jsonResp.String())
}
