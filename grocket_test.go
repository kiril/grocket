package grocket_test

import (
    "../grocket"
    "testing"
    "time"
)

func TestTimeBucketMarshaling(tests *testing.T) {
    bucket := &grocket.TimeBucket{Time: time.Now(), EventIds: []string{}}
    if bucket == nil {
        tests.Fatal("OMG")
    }
}
