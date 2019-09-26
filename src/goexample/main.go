package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query()["name"][0]
		fmt.Fprintf(w, "Hello, %q\n, your lucky number is: %d", name, generateLuckyNumber(name))
	})
	err := http.ListenAndServe(":18770", nil)
	log.Fatal(err)
}

func generateLuckyNumber(name string) int {
	return len(name)
}
