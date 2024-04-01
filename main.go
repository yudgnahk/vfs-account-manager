package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	fmt.Println("Server is running on http://localhost:8080/ping")
	http.ListenAndServe(":8080", nil)
}
