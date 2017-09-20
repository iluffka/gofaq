package main

import (
	"log"
	"net/http"
	"fmt"
	"io"
)
//example from package
type apiHandler struct {}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.FormValue("api"))
	if _, err := io.WriteString(w, "<h1>" + r.RequestURI + "</h1>"); err != nil {
		log.Fatalln("Something wrong")
	}
}

func main() {
	h := http.NewServeMux()
	h.Handle("/test/", apiHandler{})
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			NotFoundHandle(w, r)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8091", h))
}

func NotFoundHandle(w http.ResponseWriter, r *http.Request) {
	handleError(w, "error handle", 404)
}

func handleError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte("<h1>404</h1>"))
	fmt.Fprintln(w, error)
}