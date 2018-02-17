package main

import (
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	name string
}

func (p Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My name is %s\n", p.name)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url is %s\n", r.URL.Path)
}

func main() {
	mux := http.NewServeMux()

	personHandler := Person{"faizal"}

	mux.HandleFunc("/", Index)
	mux.Handle("/about", personHandler)

	files := http.FileServer(http.Dir("/Users/mohammadfaizal/go_workspace/src/github.com/zeenfaizpy/GoTodo"))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	server := http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
