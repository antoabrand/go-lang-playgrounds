package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Dockerizing my go web app")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Web World")
	})

	//localhost:8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}
