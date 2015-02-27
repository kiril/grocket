package main

import (
	"fmt"
	"log"
	"time"
	"net/http"

	"github.com/gorilla/mux"
)

var version = "0.0.1"

type Event struct {
	Time        time.Time `json:"due"`
	Payload     string    `json:"payload"`
	Expiry      time.Time `json:"expiry"`
	EndPoint    string    `json:"endpoint"`
	MaxAttempts int       `json:"max-attempts"`
	Verb        string    `json:"verb"`
}

func main() {
	fmt.Printf("Grocket Server version %s\n", version);
	fmt.Printf("starting up...")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", StatusInfo)
	router.HandleFunc("/events", ScheduleEvent)
	router.HandleFunc("/events/{eventId}", ViewEvent)
	router.HandleFunc("/events/{eventId}/due", ViewEventDue)
	router.HandleFunc("/events/{eventId}/expiry", ViewEventExpiry)
	router.HandleFunc("/events/{eventId}/payload", ViewEventPayload)
	router.HandleFunc("/events/{eventId}/endpoint", ViewEventEndpoint)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func StatusInfo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Grocket Server version %s", version)
}

func ScheduleEvent(writer http.ResponseWriter, request *http.Request) {
}

func ViewEvent(writer http.ResponseWriter, request *http.Request) {
}

func ViewEventDue(writer http.ResponseWriter, request *http.Request) {
}

func ViewEventExpiry(writer http.ResponseWriter, request *http.Request) {
}

func ViewEventPayload(writer http.ResponseWriter, request *http.Request) {
}


func ViewEventEndpoint(writer http.ResponseWriter, request *http.Request) {
}
