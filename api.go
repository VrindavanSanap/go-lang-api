package main

import (
	"fmt"
	"net"
	"net/http"
	"log"
)

// homeHandler handles requests to the root path "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate the response is plain text
	w.Header().Set("Content-Type", "text/plain")
    	fmt.Printf("served a request\n")
	// Write the response to the client
	fmt.Fprintf(w, "Hello, World! This is a simple Go API. \n")
ipAndPort := r.RemoteAddr
	ip, _, err := net.SplitHostPort(ipAndPort)
    if err != nil {
        // Handle error, e.g., if the format is invalid
        ip = ipAndPort // Fallback
    }

    fmt.Printf("IP (RemoteAddr) is: %s\n", ip)


}

func main() {
	// 1. Define a route and associate it with a handler function
	// http.HandleFunc registers the handler for the given pattern in the DefaultServeMux.
	http.HandleFunc("/", homeHandler)

	// 2. Start the server
	port := ":8080"
	fmt.Printf("Starting server on http://localhost%s\n", port)
	
	// http.ListenAndServe starts an HTTP server with a given address and handler.
	// nil means use the DefaultServeMux.
	// log.Fatal will print the error and exit if the server fails to start (e.g., port already in use).
	log.Fatal(http.ListenAndServe(port, nil))
}
