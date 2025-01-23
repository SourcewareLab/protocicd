package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Setting up mux")
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Creating index route")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})

	fmt.Println("Server listening on :8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error:", err)
	}
}
