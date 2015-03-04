package main

import (
    client "github.com/kiril/grocket/client"
    "os"
	"fmt"
	"log"
    "strconv"
	"net/http"
)

var version = "0.0.2"

func StartServer(port int) {
	fmt.Printf("Grocket Server starting up on port %d...\n", port)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func main() {
	fmt.Printf("Grocket version %s\n", version);

    args := os.Args[1:]
    command := "server"
    if len(args) >= 1 {
        command = args[0]
    }

    switch command {
    case "server":
        port := 8080
        if len(args) > 1 {
            parsedPort, error := strconv.Atoi(args[1])
            if error != nil {
                log.Fatalf("Invalid server port: needs to be int, not %s", args[1])
            }
            port = parsedPort
        }
        StartServer(port)

    case "client":
        port, error := strconv.Atoi(args[1])
        if error != nil {
            log.Fatalf("Invalid server port: needs to be int, not %s", args[1])
        }
        client.StartClient(port)
    }
}
