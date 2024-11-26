package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "hello from golang")
}

func main() {
	http.HandleFunc("/api/test", helloHandler)

	port := "8080"
	fmt.Printf("Server is running on http://localhost:%s/api/test\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
