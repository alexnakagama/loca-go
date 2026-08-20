package main

import (
	"fmt"
	"net/http"
)

func main() {

	// When a request to / execute this function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Loca-Go API")
	})

	fmt.Println("server running on port: 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
