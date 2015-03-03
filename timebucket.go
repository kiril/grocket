package grocket

import (
    "time"
    "fmt"
    "log"
)


type TimeBucket struct {
    Time     time.Time
    EventIds []string
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

