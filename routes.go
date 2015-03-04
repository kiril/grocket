package main

import (
	"net/http"
)


type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Status",      "GET",  "/",                                  StatusInfo,},

	Route{"Due",         "GET",  "/events/{eventId}/due",              EventDue,},
	Route{"Expiry",      "GET",  "/events/{eventId}/expiry",           EventExpiry,},
	Route{"Payload",     "GET",  "/events/{eventId}/payload",          EventPayload,},
	Route{"Endpoint",    "GET",  "/events/{eventId}/endpoint",         EventEndpoint,},
	Route{"MaxAttempts", "GET",  "/events/{eventId}/max-attempts",     EventMaxAttempts,},
	Route{"Verb",        "GET",  "/events/{eventId}/verb",             EventVerb,},
	Route{"Due",         "GET",  "/events/{eventId}",                  ViewEvent,},

	Route{"Events",      "POST", "/events",                            ScheduleEvent,},

	Route{"Events",      "GET",  "/events/list/{startTime}-{endTime}", EventList,},
}
