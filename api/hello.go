package handler

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Name		string
}


func HelloWorld(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Johnni"}
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
