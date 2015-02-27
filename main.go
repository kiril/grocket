package main

import "fmt"
import "net/http"

var version = "0.0.1"

func main() {
	fmt.Printf("Grocket Server version %s\n", version);
	fmt.Printf("starting up...")
	http.HandleFunc("/", basicInfo)
	http.ListenAndServe(":8080", nil)
	fmt.Printf("shutting down...")
}

func basicInfo(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte(fmt.Sprintf("Grocket Server version %s", version)))
}
