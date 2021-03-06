package main

import (
    "time"
    "log"

    core "github.com/kiril/grocket/core"
    btree "github.com/kiril/btree"
)

type IndexedEvent struct {
    Event  *core.Event
    Bucket *TimeBucket
}

var eventsById = make(map[string]*IndexedEvent)
var bucketByTimeIndex = btree.NewBtree()

func FindBucketByTime(due time.Time) *TimeBucket {
    key, error := due.MarshalBinary()
    if error != nil {
        log.Fatal(error)
    }
    binary, searchError := bucketByTimeIndex.Search(key)
    if error != nil {
        log.Fatal(searchError)
    }
    if binary == nil {
        return nil
    }
    bucket := &TimeBucket{}
    bucket.UnmarshalBinary(binary)
    return bucket

}

func FindOrCreateTimeBucket(due time.Time) *TimeBucket {
    bucket := FindBucketByTime(due)
    if bucket == nil {
        bucket = &TimeBucket{Time: due, EventIds: [][]byte{}}
        SaveTimeBucket(bucket)
    }
    return bucket
}

func NextTimeBucket() *TimeBucket {
    first, error := bucketByTimeIndex.Left()
    if error != nil {
        return nil
    }

    if first != nil {
        bucket := &TimeBucket{}
        bucket.UnmarshalBinary(first)
        return bucket
    }

    return nil
}

func SaveTimeBucket(bucket *TimeBucket) error {
    key, timeError := bucket.Time.MarshalBinary()
    if timeError != nil {
        log.Fatal(timeError)
    }

    value, valueError := bucket.MarshalBinary()
    if valueError != nil {
        log.Fatal(valueError)
    }

    existing := FindBucketByTime(bucket.Time)
    if existing != nil {
        return bucketByTimeIndex.Update(key, value)
    } else {
        return bucketByTimeIndex.Insert(key, value)
    }
}

func RemoveTimeBucket(bucket *TimeBucket) error {
    key, timeError := bucket.Time.MarshalBinary()
    if timeError != nil {
        log.Fatal(timeError, "couldn't marshal time")
    }
    return bucketByTimeIndex.Delete(key)
}

func (indexed IndexedEvent) AddToBucket() {
    indexed.Bucket.AddEvent(indexed.Event)
}

func StoreEvent(event *core.Event) {
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

    indexed.AddToBucket()
    SaveTimeBucket(indexed.Bucket)
    eventsById[event.Id] = indexed
}

func RetrieveEventById(id string) *core.Event {
    indexed := eventsById[id]
    if indexed == nil {
        return nil
    } else {
        return indexed.Event
    }
}

func ClearEvent(id string) {
    indexed := eventsById[id]
    if indexed != nil {
        bucket := FindBucketByTime(indexed.Event.Due)
        bucket.RemoveEvent(indexed.Event)
        if bucket.IsEmpty() {
            RemoveTimeBucket(bucket)
        }
        delete(eventsById, id)
    }
}

func CountBuckets() int {
    count, error := bucketByTimeIndex.Count()
    if error != nil {
        log.Fatal(error)
    }
    return count
}

func ProbabilisticSleepDuration() time.Duration {
    return time.Millisecond * 100
}
