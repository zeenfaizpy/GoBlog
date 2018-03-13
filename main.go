package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/zeenfaizpy/GoBlog/src/config"
	"github.com/zeenfaizpy/GoBlog/src/views"
)

func main() {

	db, err := config.NewDB("postgres://bloguser:blogpassword@localhost/blog?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	env := config.Env{DB: db}

	mux := mux.NewRouter()
	mux.HandleFunc("/", views.IndexHandler(env))

	server := http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: mux,
	}
	log.Print("Running Server at http://0.0.0.0:8000")

	log.Fatal(server.ListenAndServe())

}
