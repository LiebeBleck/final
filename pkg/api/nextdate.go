package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func afterNow(now, date time.Time) bool {
	dateYear, dateMonth, dateDay := date.Date()
	nowYear, nowMonth, nowDay := now.Date()
	dateTruncated := time.Date(dateYear, dateMonth, dateDay, 0, 0, 0, 0, time.UTC)
	nowTruncated := time.Date(nowYear, nowMonth, nowDay, 0, 0, 0, 0, time.UTC)
	return dateTruncated.After(nowTruncated)
}

func NextDate(now time.Time, dstart, repeat string) (string, error) {
	if repeat == "" {
		return "", fmt.Errorf("repeat rule is empty")
	}

	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %v", err)
	}

	parts := strings.Split(repeat, " ")
	if len(parts) < 2 && parts[0] != "y" {
		return "", fmt.Errorf("invalid repeat rule format")
	}

	var next time.Time
	switch parts[0] {
	case "d":
		days, err := strconv.Atoi(parts[1])
		if err != nil || days <= 0 {
			return "", fmt.Errorf("invalid days in repeat rule: %v", err)
		}
		next = date
		for !afterNow(now, next) {
			next = next.AddDate(0, 0, days)
		}
	case "y":
		next = date
		for !afterNow(now, next) {
			next = next.AddDate(1, 0, 0)
		}
	default:
		return "", fmt.Errorf("unsupported repeat rule: %s", parts[0])
	}

	return next.Format(DateFormat), nil
}

func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJson(w, map[string]string{"error": "method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	nowStr := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	if nowStr == "" || date == "" {
		writeJson(w, map[string]string{"error": "now and date parameters are required"}, http.StatusBadRequest)
		return
	}

	now, err := time.Parse(DateFormat, nowStr)
	if err != nil {
		writeJson(w, map[string]string{"error": fmt.Sprintf("invalid now format: %v", err)}, http.StatusBadRequest)
		return
	}

	next, err := NextDate(now, date, repeat)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	writeJson(w, map[string]string{"date": next}, http.StatusOK)
}
