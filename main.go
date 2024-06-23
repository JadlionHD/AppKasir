package main

import (
	"log"
	"net/http"

	"github.com/JadlionHD/AppKasir/internal/resources"
	"github.com/JadlionHD/AppKasir/internal/routes"
)

func main() {

	assets, _ := resources.Assets()

	fs := http.FileServer(http.FS(assets))
	http.Handle("/", fs)

	http.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", "Hello world")
	})
	http.HandleFunc("/api/login", routes.HandleLogin)

	log.Println("Running on http://localhost:80")
	http.ListenAndServe("localhost:80", nil)

}
