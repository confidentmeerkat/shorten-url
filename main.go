package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Web!\n")
}
