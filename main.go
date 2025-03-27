package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("GO")

	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(writer, "Welcome tot Go2, the URL Shortener.")
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
