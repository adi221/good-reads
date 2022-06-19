package api

import (
	"encoding/json"
	"net/http"
)

type InfoConfig struct {
	Author  string `json:"author"`
	Version string `json:"version"`
}

func index() http.Handler {
	info := InfoConfig{
		Author:  "Adi Mizrahi",
		Version: "v1",
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
}
