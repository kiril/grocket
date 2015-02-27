package main

import (
    "time"
    "sort"
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
    Event  *Event
    Bucket *TimeBucket
}

var eventsById map[string]*IndexedEvent
var timeBuckets list.List

func (bucket TimeBucket) RemoveEvent(event *Event) {
}

func (bucket TimeBucket) AddEvent(event *Event) {
}

func (indexed IndexedEvent) EnsureInBucket() {
    i := sort.Search(len(indexed.Bucket.EventIds),
        func(i int) bool {return indexed.Bucket.EventIds[i] >= indexed.Event.Id})

    if i >= len(indexed.Bucket.EventIds) { // all event ids are < me
        indexed.Bucket.EventIds = append(indexed.Bucket.EventIds, indexed.Event.Id)

    } else if i == 0 && indexed.Bucket.EventIds[i] != indexed.Event.Id { // greater than me
        indexed.Bucket.EventIds = append([]string{indexed.Event.Id,}, indexed.Bucket.EventIds...)

    } else if indexed.Bucket.EventIds[i] != indexed.Event.Id {
        eventIds := make([]string, len(indexed.Bucket.EventIds)+1)
        for j := 0; j < i; j++ {
            eventIds[j] = indexed.Bucket.EventIds[j]
        }

        eventIds[i] = indexed.Event.Id

        for k := i+1; k < len(eventIds); k++ {
            eventIds[k] = indexed.Bucket.EventIds[k-1]
        }

        indexed.Bucket.EventIds = eventIds
    }
}

func RemoveTimeBucket(bucket *TimeBucket) {
}

func RemoveFromTimeBucket(indexed *IndexedEvent) {
    if len(indexed.Bucket.EventIds) == 1 {
        RemoveTimeBucket(indexed.Bucket)
    }
}

func StoreEvent(event *Event) {
    indexed := eventsById[event.Id]
    if indexed != nil {
        if ! indexed.Event.Due.Equal(event.Due) {
            indexed.Bucket.RemoveEvent(indexed.Event)
            indexed = &IndexedEvent{
                Event: event,
                Bucket: FindOrCreateTimeBucket(event.Due),
            }

        } else {
            indexed.Event = event
        }

    } else {
        indexed = &IndexedEvent{
            Event: event,
            Bucket: FindOrCreateTimeBucket(event.Due),
        }
    }

    indexed.EnsureInBucket()
    eventsById[event.Id] = indexed
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
