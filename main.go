package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Server running on port 8080")

	http.Handle("/",

		http.HandlerFunc(),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
