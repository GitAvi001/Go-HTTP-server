package main


import (
	"fmt"
	"log"
	"net/http"
)
func main() {

	fmt.Println("Listening to port 8080")

	log.Fatal(http.ListenAndServe(":8080",nil))

}
