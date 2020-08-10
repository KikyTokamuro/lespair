package main

import (
	"strconv"
	"strings"
	"time"
)

func getWeek(t time.Time) int {
	_, week := t.ISOWeek()

	return week - 1
}

func getStarOfWeek(num string) string {
	i, _ := strconv.Atoi(num)

	if (i % 2) == 0 {
		return "S1"
	}

	return "S2"
}

func getStarAndDay(day string) (string, string) {
	today := time.Now()
	yesterday := today.Add(-24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	switch day {
	case "yesterday":
		d := getStarOfWeek(yesterday.Weekday().String())
		w := strings.ToLower(yesterday.Weekday().String())
		return d, w

	case "today":
		d := getStarOfWeek(today.Weekday().String())
		w := strings.ToLower(today.Weekday().String())
		return d, w

	case "tomorrow":
		d := getStarOfWeek(tomorrow.Weekday().String())
		w := strings.ToLower(tomorrow.Weekday().String())
		return d, w

	default:
		return "", ""
	}
}
