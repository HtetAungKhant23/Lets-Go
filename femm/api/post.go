package api

import (
	"encoding/json"
	"github.com/htetaungkhant/go/museum/data"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newData data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&newData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			data.Add(newData)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("OK"))
		}
	} else {
		http.Error(w, "Unsupported method", http.StatusBadRequest)
	}
}
