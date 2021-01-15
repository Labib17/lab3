package main

import (
	"flag"
	"log"
)

//PORT to connect
var PORT = flag.Int("p", 8080, "HTTP port number")

func main() {
	// Create the server.
	if err := starting_the_server(); err == nil {
		log.Println("Starting chat server...")
	} else {
		log.Fatalf("Cannot initialize  server: %s", err)
	}
}
