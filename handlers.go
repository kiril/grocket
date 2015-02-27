package main

import (
	"fmt"
    "time"
	"net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)


func StatusInfo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Grocket Server version %s", version)
}

func ScheduleEvent(writer http.ResponseWriter, request *http.Request) {
}

func ViewEvent(writer http.ResponseWriter, request *http.Request) {
    vars := mux.Vars(request)
    eventId := vars["eventId"]

    event := Event{
        Id: eventId,
        Time: time.Now(),
        Payload: "somedata",
        Expiry: time.Now().Add(time.Minute * 1),
        EndPoint: "http://dev/null",
        MaxAttempts: 0,
        Verb: "PUT",
    }

    writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
    writer.WriteHeader(http.StatusOK)

    if error := json.NewEncoder(writer).Encode(event); error != nil {
        panic(error)
    }
}

func EventDue(writer http.ResponseWriter, request *http.Request) {
}

func EventExpiry(writer http.ResponseWriter, request *http.Request) {
}

func EventPayload(writer http.ResponseWriter, request *http.Request) {
}

func EventEndpoint(writer http.ResponseWriter, request *http.Request) {
}

func EventVerb(writer http.ResponseWriter, request *http.Request) {
}

func EventMaxAttempts(writer http.ResponseWriter, request *http.Request) {
}

func EventList(writer http.ResponseWriter, request *http.Request) {
}
