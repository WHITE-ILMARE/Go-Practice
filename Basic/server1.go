// a simple web server
package main

import (
	"./lissajous"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler1(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
}
