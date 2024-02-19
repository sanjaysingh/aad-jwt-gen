package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sanjaysingh/aad-jwt-gen/cmd"
	"github.com/sanjaysingh/aad-jwt-gen/handlers"
)

func main() {
	headless := flag.Bool("headless", false, "Run in headless mode for CLI token generation")
	clientId := flag.String("clientId", "", "Client ID")
	tenantId := flag.String("tenantId", "", "Tenant ID")
	scope := flag.String("scope", "", "Scope")
	clientSecret := flag.String("clientSecret", "", "Client Secret")

	flag.Parse()

	if *headless {
		if *scope == "" {
			*scope = ".default"
		}
		if *clientId == "" || *tenantId == "" || *clientSecret == "" {
			fmt.Println("Client ID, Tenant ID, and Client Secret are required in headless mode.")
			return
		}
		cmd.RunHeadless(*clientId, *tenantId, *clientSecret, *scope)
	} else {
		// Serve the HTML file
		http.Handle("/", http.FileServer(http.Dir("./static")))

		// Handle JWT token generation
		http.HandleFunc("/generate-token", handlers.TokenHandler)

		// Start the web server
		log.Println("Starting server on :5555...")
		log.Fatal(http.ListenAndServe(":5555", nil))
	}
}
