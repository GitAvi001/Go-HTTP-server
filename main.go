package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")

}

func main() {

	fmt.Println("Server running on port 8080")

	http.Handle("/",

		http.HandlerFunc(),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
