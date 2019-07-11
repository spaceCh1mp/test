package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleIndex)
	log.Println(http.ListenAndServe(":8000", nil))
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
