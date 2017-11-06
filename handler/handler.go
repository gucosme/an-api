package handler

import (
	"encoding/json"
	"github.com/gucosme/an-api/model"
	"net/http"
	"strconv"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	age, _ := strconv.Atoi(query.Get("age"))

	p := &model.Person{Name: query.Get("name"), Age: age}

	per, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(per)
}
