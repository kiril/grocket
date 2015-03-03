package grocket_test

import (
    "../grocket"
    "testing"
    "time"
    "reflect"
)

func TestTimeBucketMarshaling(tests *testing.T) {
    bucket := &grocket.TimeBucket{Time: time.Now(), EventIds: [][]byte{[]byte("111")}}

    bytes, error := bucket.MarshalBinary()
    if error != nil {
        tests.Fatal(error)
    }

    bucket2 := &grocket.TimeBucket{}
    bucket2.UnmarshalBinary(bytes)

    if ! reflect.DeepEqual(bucket, bucket2) {
        tests.Fatal("Failed to round-trip")
    }
}

func TestBucketRoundTrips(tests *testing.T) {
    bucket := &grocket.TimeBucket{Time: time.Now(), EventIds: [][]byte{[]byte("111")}}
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
