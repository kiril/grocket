package grocket

import (
    "time"
    "sort"
    "log"
    "fmt"

    "github.com/kiril/btree"
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

func (bucket *TimeBucket) String() string {
    return fmt.Sprintf("{Event:%s, EventIds:%s}", bucket.Time, bucket.EventIds)
}

func (bucket TimeBucket) MarshalBinary() ([]byte, error) {
    timeBytes, error := bucket.Time.MarshalBinary()
    if error != nil {
        log.Fatal(error)
    }

    idCount := len(bucket.EventIds)

    idByteStrings := make([][]byte, idCount)

    totalByteLength := 1 + len(timeBytes) + 1 + idCount

    for i := 0; i < idCount; i++ {
        idByteStrings[i] = []byte(bucket.EventIds[i])
        totalByteLength += len(idByteStrings[i])
    }

    bytes := make([]byte, totalByteLength)

    b := 0
    bytes[b] = byte(len(timeBytes))
    b++

    for i := 0; i < len(timeBytes); i++ {
        bytes[b] = timeBytes[i]
        b++
    }

    bytes[b] = byte(idCount)
    b++

    for i := 0; i < idCount; i++ {
        bytes[b] = byte(len(idByteStrings[i]))
        b++
        for j := 0; j < len(idByteStrings[i]); j++ {
            bytes[b] = idByteStrings[i][j]
            b++
        }
    }

    return bytes, nil
}

func (bucket *TimeBucket) UnmarshalBinary(bytes []byte) error {
    // 0 : length of timeBytes
    // 1-<length of timeBytes> timeBytes
    // then:
    // L : one byte length of a string id
    // L+1... a string
    lengthOfTimeBytes := int(bytes[0])
    timeBytes := make([]byte, lengthOfTimeBytes)
    for i := 0; i < lengthOfTimeBytes; i++ {
        timeBytes[i] = bytes[i+1]
    }

    due := time.Now()
    error := due.UnmarshalBinary(timeBytes)
    if error != nil {
        log.Fatal(error)
    }
    bucket.Time = due

    b := lengthOfTimeBytes + 1
    idCount := int(bytes[b])
    b++

    bucket.EventIds = make([]string, idCount)
    for i := 0; i < idCount; i++ {
        idLength := int(bytes[b])
        b++
        idBytes := make([]byte, idLength)
        for j := 0; j < idLength; j++ {
            idBytes[j] = bytes[b+j]
        }
        bucket.EventIds[i] = string(idBytes)
        b += idLength
    }

    return nil
}


var eventsById map[string]*IndexedEvent
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
    bucket := &TimeBucket{}
    bucket.UnmarshalBinary(binary)
    return bucket

}

func FindOrCreateTimeBucket(due time.Time) *TimeBucket {
    bucket := FindBucketByTime(due)
    if bucket == nil {
        bucket = &TimeBucket{Time: due, EventIds: []string{}}
        InsertTimeBucket(bucket)
    }
    return bucket
}

func NextTimeBucket() *TimeBucket {
    first, error := bucketByTimeIndex.Left()
    if error != nil {
        log.Fatal(error)
    }

    if first != nil {
        bucket := &TimeBucket{}
        bucket.UnmarshalBinary(first)
        return bucket
    }

    return nil
}

func InsertTimeBucket(bucket *TimeBucket) error {
    key, timeError := bucket.Time.MarshalBinary()
    if timeError != nil {
        log.Fatal(timeError)
    }

    value, valueError := bucket.MarshalBinary()
    if valueError != nil {
        log.Fatal(valueError)
    }

    return bucketByTimeIndex.Insert(key, value)
}

func RemoveTimeBucket(bucket *TimeBucket) error {
    key, timeError := bucket.Time.MarshalBinary()
    if timeError != nil {
        log.Fatal(timeError)
    }
    return bucketByTimeIndex.Delete(key)
}


func (bucket TimeBucket) RemoveEvent(event *Event) {
    i := sort.Search(len(bucket.EventIds),
        func(i int) bool {return bucket.EventIds[i] >= event.Id})

    if i < len(bucket.EventIds) && bucket.EventIds[i] == event.Id {
        if ( len(bucket.EventIds) == 1 ) {
            bucket.EventIds = []string{}

        } else {
            eventIds := make([]string, len(bucket.EventIds)-1)
            for j := 0; j < i; j++ {
                eventIds[j] = bucket.EventIds[j]
            }

            for k := i; k < len(eventIds); k++ {
                eventIds[k] = bucket.EventIds[k+1]
            }

            bucket.EventIds = eventIds
        }
    }
}

func (bucket TimeBucket) AddEvent(event *Event) {
    i := sort.Search(len(bucket.EventIds),
        func(i int) bool {return bucket.EventIds[i] >= event.Id})

    if i >= len(bucket.EventIds) { // all event ids are < me
        bucket.EventIds = append(bucket.EventIds, event.Id)

    } else if i == 0 && bucket.EventIds[i] != event.Id { // greater than me
        bucket.EventIds = append([]string{event.Id,}, bucket.EventIds...)

    } else if bucket.EventIds[i] != event.Id {
        eventIds := make([]string, len(bucket.EventIds)+1)
        for j := 0; j < i; j++ {
            eventIds[j] = bucket.EventIds[j]
        }

        eventIds[i] = event.Id

        for k := i+1; k < len(eventIds); k++ {
            eventIds[k] = bucket.EventIds[k-1]
        }

        bucket.EventIds = eventIds
    }
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
    eventsById[event.Id] = indexed
}

func RetrieveEventById(id string) *Event {
    return eventsById[id].Event
}

func ProbabilisticSleepDuration() time.Duration {
    return time.Millisecond * 100
}
