package grocket_test

import (
    "../grocket"
    "testing"
    "time"
)

func TestTimeBucketMarshaling(tests *testing.T) {
    bucket := &grocket.TimeBucket{Time: time.Now(), EventIds: []string{"111"}}

    bytes, error := bucket.MarshalBinary()
    if error != nil {
        tests.Fatal(error)
    }

    bucket2 := &grocket.TimeBucket{}
    bucket2.UnmarshalBinary(bytes)

    if len(bucket2.EventIds) != 1 {
        tests.Fatalf("Didn't get right number of events back (%d)", len(bucket2.EventIds))
    }

    if bucket.EventIds[0] != bucket2.EventIds[0] {
        tests.Error("Shit")
    }
}
