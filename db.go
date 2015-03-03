package grocket

import (
    "time"
    "log"

    "github.com/kiril/btree"
)

/*
 * So, let's figure out how we're going to store this shit.
 * I want an index from {id -> event}
 * and I want an ordered index from {due -> event}
 */

type IndexedEvent struct {
    Event  *Event
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

    indexed.AddToBucket()
    SaveTimeBucket(indexed.Bucket)
    eventsById[event.Id] = indexed
}

func RetrieveEventById(id string) *Event {
    return eventsById[id].Event
}

func ProbabilisticSleepDuration() time.Duration {
    return time.Millisecond * 100
}
