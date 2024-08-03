package api

import (
	"encoding/json"
	"github.com/htetaungkhant/go/museum/data"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query()["id"]
	if id != nil {
		finalId, err := strconv.Atoi(id[0])
		if err != nil {
			http.Error(w, "Invalid Exhibition", http.StatusBadRequest)
		} else {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
