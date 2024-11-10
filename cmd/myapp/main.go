package main

import (
    "fmt"
    "log"
    "net/http"
    "example.com/exampleApp/api"
)

func add(a int, b int) int {
    return a + b
}

func main() {
    // Initialize the router using the api package
    r := api.SetupRouter()

    // Register the router with the default http server
    http.Handle("/", r)

    // Start the HTTP server in a separate goroutine
    go func() {
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
            log.Fatalf("Server failed to start: %v", err)
        }
    }()

    // Simple calculation for example purposes
    var x, y int = 5, 3
    sum := add(x, y)
    fmt.Printf("Sum of x and y is: %d\n", sum)

    // Keep the program running
    select {}
}
