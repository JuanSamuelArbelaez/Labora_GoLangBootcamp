package main

import (
"fmt"
"net/http"
)

func main() {
	port := 8087

	http.HandleFunc("/", handler)

	fmt.Printf("Starting server on port %d\n", port)

	// Port Set-Up
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Servidor HTTP running!")
}