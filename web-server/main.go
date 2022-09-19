// Taken & modified from the following DigitalOcean tutorial article:
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
package main

import (
	"io"
	"log"
	"net/http"
)

// Address & port the server listens on (no leading address means "listen on all")
const addr = ":8080"

// Logic for handling GET requests on "/"
func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("got request on /")

	_, err := io.WriteString(w, "you hit: /\n")
	if err != nil {
		log.Printf("error writing response: %s\n", err.Error())
	}
}

// Logic for handling GET requests on "/healthcheck"
func getHealthcheck(w http.ResponseWriter, r *http.Request) {
	log.Println("got request on /healthcheck")

	_, err := io.WriteString(w, "you hit: /healthcheck\n")
	if err != nil {
		log.Printf("error writing response: %s\n", err.Error())
	}
}

func main() {
	// Register your functions with their URL paths
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/healthcheck", getHealthcheck)

	// Start the server
	log.Printf("Starting server on %s\n", addr)
	err := http.ListenAndServe(addr, nil) // second arg is for an explicit HTTP handler -- nil just uses the default
	if err != nil {
		log.Fatalf("error starting server: %s\n", err.Error())
	}
}
