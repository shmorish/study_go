package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc!\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
