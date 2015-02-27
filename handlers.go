package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)


func StatusInfo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Grocket Server version %s", version)
}

func ScheduleEvent(writer http.ResponseWriter, request *http.Request) {
}

func ViewEvent(writer http.ResponseWriter, request *http.Request) {
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
