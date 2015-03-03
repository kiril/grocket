package grocket

import (
    "time"
    "fmt"
    "log"
)


type TimeBucket struct {
    Time     time.Time
    EventIds [][]byte
}

func (bucket *TimeBucket) String() string {
    stringIds := make([]string, len(bucket.EventIds))
    for i := 0; i < len(stringIds); i++ {
        stringIds[i] = string(bucket.EventIds[i])
    }
    return fmt.Sprintf("{Event:%s, EventIds:%s}", bucket.Time, stringIds)
}

func (bucket TimeBucket) MarshalBinary() ([]byte, error) {
    timeBytes, error := bucket.Time.MarshalBinary()
    if error != nil {
        log.Fatal(error)
    }

    idCount := len(bucket.EventIds)

    totalByteLength := 1 + len(timeBytes) + 1 + idCount

    for i := 0; i < idCount; i++ {
        totalByteLength += len(bucket.EventIds[i])
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
        bytes[b] = byte(len(bucket.EventIds[i]))
        b++
        for j := 0; j < len(bucket.EventIds[i]); j++ {
            bytes[b] = bucket.EventIds[i][j]
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

    bucket.EventIds = make([][]byte, idCount)
    for i := 0; i < idCount; i++ {
        idLength := int(bytes[b])
        b++
        idBytes := make([]byte, idLength)
        for j := 0; j < idLength; j++ {
            idBytes[j] = bytes[b+j]
        }
        bucket.EventIds[i] = idBytes
        b += idLength
    }

    return nil
}

