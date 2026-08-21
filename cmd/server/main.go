package main

import (
	"fmt"
	"net/http"

	"github.com/alexnakagama/loca-go/interal/handler"
)

func main() {
	http.HandleFunc("/locations", handler.HandleLocation)

	fmt.Println("server running on port: 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
