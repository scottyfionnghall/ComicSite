package main

import (
	h "comicsite/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", h.MakeHandler(h.ViewHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
