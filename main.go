package grocket

import (
	"fmt"
	"log"
	"net/http"
)

var version = "0.0.1"

func main() {
	fmt.Printf("Grocket Server version %s\n", version);
	fmt.Printf("starting up...")

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
