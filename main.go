package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/sanjaysingh/aad-jwt-gen/cmd"
	"github.com/sanjaysingh/aad-jwt-gen/handlers"
)

//go:embed static
var staticFiles embed.FS

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
		var staticFS = fs.FS(staticFiles)
		htmlfs, _ := fs.Sub(staticFS, "static")
		http.Handle("/", http.FileServer(http.FS(htmlfs)))

		// Handle JWT token generation
		http.HandleFunc("/generate-token", handlers.TokenHandler)

		// Start the web server
		log.Println("Starting server on :5555...")
		log.Fatal(http.ListenAndServe(":5555", nil))
	}
}
