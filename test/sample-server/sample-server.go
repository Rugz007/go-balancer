package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	// Check if the port argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run sample-server.go <port>")
		return
	}

	// Get the port number from the command-line argument
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid port number")
		return
	}

	// Create a request handler
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		// Sleep for 25 milliseconds to simulate a slow response
		time.Sleep(25 * time.Millisecond)
		fmt.Fprintf(ctx, "Hello, World!")
	}

	// Start the server on the specified port
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), requestHandler); err != nil {
		fmt.Printf("Error in ListenAndServe: %s\n", err)
	}
}
