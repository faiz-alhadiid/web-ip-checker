package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/check-ip", checkIP)

	log.Println("Server running at port " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("Server stopped")
	}

}
