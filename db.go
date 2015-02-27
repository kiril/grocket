package main

import "time"

/*
 * So, let's figure out how we're going to store this shit.
 * I want an index from {id -> event}
 * and I want an ordered index from {due -> event}
 */

func StoreEvent(event Event) {
}

func RetrieveEventById(id string) {
}

func ProbabilisticSleepDuration() time.Duration {
    return time.Millisecond * 100
}

func NextDueEvent() Event {
    return Event{
        Id:          "id",
        Time:        time.Now(),
        Payload:     "somedata",
        Expiry:      time.Now().Add(time.Minute * 1),
        EndPoint:    "http://dev/null",
        MaxAttempts: 0,
        Verb:        "PUT",
    }
}
