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

type IndexedEvent struct {
    Event *Event
    Time  time.Time
}

var eventsById map[string]*Event
var timeBuckets list.List

func RemoveFromTimeBucket(event *Event) {
}

func AddToTimeBucket(event *Event) {
}

func StoreEvent(event *Event) {
    indexedEvent := eventsById[event.Id]
    if indexedEvent != nil {
        if ! indexedEvent.Due.Equal(event.Due) {
            RemoveFromTimeBucket(indexedEvent)
            AddToTimeBucket(event)
        }
    } else {
        AddToTimeBucket(event)
    }

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
