package main

import (
    "time"
    "container/list"
)

/*
 * So, let's figure out how we're going to store this shit.
 * I want an index from {id -> event}
 * and I want an ordered index from {due -> event}
 */

type TimeBucket struct {
    Time     time.Time
    EventIds []string
}

var eventsById map[string]*Event
var timeBuckets list.List
// timeBuckets.PushBack(v)
// l.Front()

func StoreEvent(event *Event) {
    eventsById[event.Id] = event
}

func RetrieveEventById(id string) *Event {
    return eventsById[id]
}

func ProbabilisticSleepDuration() time.Duration {
    return time.Millisecond * 100
}

func NextTimeBucket() *TimeBucket {
    element := timeBuckets.Front()
    if element != nil {
        return element.Value.(*TimeBucket)
    }
    return nil
}

func RemoveTimeBucket(bucket *TimeBucket) {
}
