package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func checkIP(w http.ResponseWriter, r *http.Request) {
	endpoint := r.URL.Query().Get("url")

	temp, err := getDigResult(endpoint)

	if err != nil {
		w.WriteHeader(502)
		b, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.Write(b)
		return
	}

	result := parseDigResult(temp)

	b, _ := json.Marshal(result)

	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func index(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("static/index.html")
	if err != nil {
		w.WriteHeader(502)
		w.Write([]byte("Internal Server Error"))
		return
	}

	io.Copy(w, f)
}
