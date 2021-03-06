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
	Route{"Status",   "GET",  "/",                                  StatusInfo,},

	Route{"Next",     "GET",  "/events/next",                       NextEvent,},

	Route{"Events",   "GET",  "/events/list/{startTime}-{endTime}", EventList,},

	Route{"Event",    "GET",  "/events/{eventId}",                  EventById,},

	Route{"Schedule", "POST", "/events",                            ScheduleEvent,},
}
