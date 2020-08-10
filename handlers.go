package main

import (
	"fmt"
	"net/http"

	"github.com/tidwall/gjson"
)

// IndexPageData struct for storage data for index page
type IndexPageData struct {
	GroupOneYesterday,
	GroupOneToday,
	GroupOneTomorrow,
	GroupTwoYesterday,
	GroupTwoToday,
	GroupTwoTomorrow []string
}

// Index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := IndexPageData{}

	starYesterday, dayYesterday := getStarAndDay("yesterday")
	starToday, dayToday := getStarAndDay("today")
	starTomorrow, dayTomorrow := getStarAndDay("tomorrow")

	// GroupOneYesterday
	result := gjson.Get(string(db), fmt.Sprintf("group1.%s.%s", dayYesterday, starYesterday))
	result.ForEach(func(k, v gjson.Result) bool {
		data.GroupOneYesterday = append(data.GroupOneYesterday, v.String())
		return true
	})

	// GroupOneToday
	result = gjson.Get(string(db), fmt.Sprintf("group1.%s.%s", dayToday, starToday))
	result.ForEach(func(k, v gjson.Result) bool {
		data.GroupOneToday = append(data.GroupOneToday, v.String())
		return true
	})

	// GroupOneTomorrow
	result = gjson.Get(string(db), fmt.Sprintf("group1.%s.%s", dayTomorrow, starTomorrow))
	result.ForEach(func(k, v gjson.Result) bool {
		data.GroupOneTomorrow = append(data.GroupOneTomorrow, v.String())
		return true
	})

	// GroupTwoYesterday
	result = gjson.Get(string(db), fmt.Sprintf("group2.%s.%s", dayYesterday, starYesterday))
	result.ForEach(func(k, v gjson.Result) bool {
		data.GroupTwoYesterday = append(data.GroupTwoYesterday, v.String())
		return true
	})

	// GroupTwoToday
	result = gjson.Get(string(db), fmt.Sprintf("group2.%s.%s", dayToday, starToday))
	result.ForEach(func(k, v gjson.Result) bool {
		data.GroupTwoToday = append(data.GroupTwoToday, v.String())
		return true
	})

	// GroupTwoTomorrow
	result = gjson.Get(string(db), fmt.Sprintf("group2.%s.%s", dayTomorrow, starTomorrow))
	result.ForEach(func(k, v gjson.Result) bool {
		data.GroupTwoTomorrow = append(data.GroupTwoTomorrow, v.String())
		return true
	})

	indexTmpl.Execute(w, data)
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
