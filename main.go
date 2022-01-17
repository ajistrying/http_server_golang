package main

import (
	"log"
	"net/http"
	"time"
)

const addr = "localhost:8080"

func main() {
	//A ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.
	sm := http.NewServeMux();

	// Have the ServeMux use the handler func we created to respond to the root path "/"
	sm.HandleFunc("/", testHandler)

	// Create a server that takes in a ServMux, an address, and read/write timeouts 
	server := http.Server{
		Handler: sm,
		Addr: addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	// Tell the server to listen and serve our http handler
	err := server.ListenAndServe()

	log.Fatal(err)

}

// This handler adds a Content-Type header, a status code of 200, and an empty JSON body to the response
func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}


