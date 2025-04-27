package api

import (
	"fmt"
	"go1f/pkg/db"
	"net/http"
	"time"
)

func taskDoneHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJson(w, map[string]string{"error": "method not allowed"}, http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		writeJson(w, map[string]string{"error": "id is required"}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	if task.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			writeJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
			return
		}
	} else {

		now := time.Now()
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			writeJson(w, map[string]string{"error": fmt.Sprintf("failed to calculate next date: %v", err)}, http.StatusBadRequest)
			return
		}

		task.Date = next
		if err := db.UpdateTask(task); err != nil {
			writeJson(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
			return
		}
	}

	writeJson(w, map[string]interface{}{}, http.StatusOK)
}
