package main

import (
	"strconv"
	"strings"
	"time"
)

func getWeek(t time.Time) int {
	_, week := t.ISOWeek()

	return week
}

func getStarOfWeek(num string) string {
	i, _ := strconv.Atoi(num)

	if (i % 2) == 0 {
		return "S2"
	}

	return "S1"
}

func getStarAndDay(day string) (string, string) {
	today := time.Now()
	yesterday := today.Add(-24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	switch day {
	case "yesterday":
		d := getStarOfWeek(string(getWeek(yesterday)))
		w := strings.ToLower(yesterday.Weekday().String())
		return d, w

	case "today":
		d := getStarOfWeek(string(getWeek(yesterday)))
		w := strings.ToLower(today.Weekday().String())
		return d, w

	case "tomorrow":
		d := getStarOfWeek(string(getWeek(yesterday)))
		w := strings.ToLower(tomorrow.Weekday().String())
		return d, w

	default:
		return "", ""
	}
}
