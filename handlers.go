package main

import (
	"fmt"
	"net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)


func StatusInfo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Grocket Server version %s, %d buckets indexed", version, CountBuckets())
}

func ScheduleEvent(writer http.ResponseWriter, request *http.Request) {
}

func EventById(writer http.ResponseWriter, request *http.Request) {
    vars := mux.Vars(request)
    eventId := vars["eventId"]

    event := RetrieveEventById(eventId)

    if event == nil {
        writer.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(writer, "Event %s not found", eventId)

    } else {
        writer.Header().Set("Content-Type", "application/json;charset=UTF-8")
        writer.WriteHeader(http.StatusOK)

        if error := json.NewEncoder(writer).Encode(event); error != nil {
            panic(error)
        }
    }
}

func NextEvent(writer http.ResponseWriter, request *http.Request) {
}

func EventList(writer http.ResponseWriter, request *http.Request) {
}
