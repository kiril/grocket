package main

import (
	"net/http"

	"github.com/gorilla/mux"
)


type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{"Status",      "GET",  "/",                              StatusInfo,},

	Route{"Due",         "GET",  "/events/{eventId}/due",          EventDue,},
	Route{"Expiry",      "GET",  "/events/{eventId}/expiry",       EventExpiry,},
	Route{"Payload",     "GET",  "/events/{eventId}/payload",      EventPayload,},
	Route{"Endpoint",    "GET",  "/events/{eventId}/endpoint",     EventEndpoint,},
	Route{"MaxAttempts", "GET",  "/events/{eventId}/max-attempts", EventMaxAttempts,},
	Route{"Verb",        "GET",  "/events/{eventId}/verb",         EventVerb,},
	Route{"Due",         "GET",  "/events/{eventId}",              ViewEvent,},

	Route{"Events",      "POST", "/events",                        ScheduleEvent,},
}
