package main

import (
	"log"
	"net/http"
)

var staticpath = "/build/"
var port = ":8080"

func main() {
	// http.Handle("/favicon.ico", http.FileServer(http.Dir(staticpath+"favicon.ico")))
	http.Handle("/", http.FileServer(http.Dir(staticpath)))
	log.Fatal(http.ListenAndServe(port, nil))
}
