package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! World!!!"))
	})

	log.Fatal(http.ListenAndServe(":8092", nil))
}