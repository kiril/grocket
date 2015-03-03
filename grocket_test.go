package grocket_test

import (
    "../grocket"
    "testing"
    "time"
    "reflect"
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

    if bucket.Time != bucket2.Time {
        tests.Error("Times don't match")
    }
}

func TestBucketRoundTrips(tests *testing.T) {
    bucket := &grocket.TimeBucket{Time: time.Now(), EventIds: []string{"111"}}
    grocket.InsertTimeBucket(bucket)

    if time.Now() == bucket.Time {
        tests.Fatal("Well that sucks, we're too fast to function")
    }

    bucket2 := grocket.NextTimeBucket()

    if ! reflect.DeepEqual(bucket, bucket2) {
        tests.Fatalf("NOT EQUAL %s != %s", bucket, bucket2)
    }

    bucket3 := grocket.FindBucketByTime(bucket.Time)

    if ! reflect.DeepEqual(bucket, bucket3) {
        tests.Fatalf("NOT EQUAL %s != %s", bucket, bucket3)
    }

    error := grocket.RemoveTimeBucket(bucket)
    if error != nil {
        tests.Fatal(error, "couldn't remove bucket")
    }

    bucket4 := grocket.NextTimeBucket()
    if bucket4 != nil {
        tests.Fatal("Delete didn't work")
    }
}
