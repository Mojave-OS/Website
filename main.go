package main

import (
	"fmt"
	"net/http"
	"os"
)

type GeneralServer struct {

}

func main() {

	// Initial Variable Configuration
	Port := os.Getenv("PORT")
	Gs := GeneralServer{}

	// The Mux for the Server
	mux := http.NewServeMux()
	mux.HandleFunc("/", Gs.ServeIndex)
	
	// Driver Code
	http.ListenAndServe(fmt.Sprintf(":%s", Port), mux)
}