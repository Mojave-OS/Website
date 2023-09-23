package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Initial Variable Configuration
	Port := os.Getenv("PORT")

	// The Mux for the Server
	mux := http.NewServeMux()
	

	// Driver Code
	http.ListenAndServe(fmt.Sprintf(":%s", Port), mux)
}