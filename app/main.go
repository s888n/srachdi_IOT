package main

import (
	"net/http"
	"encoding/json"
)

type PodStatus struct {
	Status      string `json:"status"`
	Version string `json:"version"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
	if err := r.Context().Err(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	data := PodStatus{
		Status: "ok",
		Version: "v2",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/healthz", healthHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	port := ":8888"
	http.ListenAndServe(port, nil)
}
