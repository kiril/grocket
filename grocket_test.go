package main_test

import (
    "../grocket"
    core "github.com/kiril/grocket/core"
    "testing"
    "time"
    "reflect"
)

func TestTimeBucketMarshaling(tests *testing.T) {
    main.CountBuckets()

    bucket := &main.TimeBucket{Time: time.Now(), EventIds: [][]byte{[]byte("111")}}

    bytes, error := bucket.MarshalBinary()
    if error != nil {
        tests.Fatal(error)
    }

    bucket2 := &main.TimeBucket{}
    bucket2.UnmarshalBinary(bytes)

    if ! reflect.DeepEqual(bucket, bucket2) {
        tests.Fatal("Failed to round-trip")
    }
}

func TestBucketRoundTrips(tests *testing.T) {
    main.FindBucketByTime(time.Now()) // tests that it's OK to be empty

    bucket := &main.TimeBucket{Time: time.Now(), EventIds: [][]byte{[]byte("111")}}
    main.SaveTimeBucket(bucket)

    if time.Now() == bucket.Time {
        tests.Fatal("Well that sucks, we're too fast to function")
    }

    bucket2 := main.NextTimeBucket()

    if ! reflect.DeepEqual(bucket, bucket2) {
        tests.Fatalf("NOT EQUAL %s != %s", bucket, bucket2)
    }

    bucket3 := main.FindBucketByTime(bucket.Time)

    if ! reflect.DeepEqual(bucket, bucket3) {
        tests.Fatalf("NOT EQUAL %s != %s", bucket, bucket3)
    }

    error := main.RemoveTimeBucket(bucket)
    if error != nil {
        tests.Fatal(error, "couldn't remove bucket")
    }

    bucket4 := main.NextTimeBucket()
    if bucket4 != nil {
        tests.Fatal("Delete didn't work")
    }
}

func TestBucketModifiers(tests *testing.T) {
    event := &core.Event{
        Id: "123",
        Due: time.Now(),
        Payload: "Holy Shit",
        Expiry: time.Now().Add(time.Second * 60),
        EndPoint: "http://gc.com/fooooo",
    }

    bucket := &main.TimeBucket{Time: time.Now(), EventIds: [][]byte{}}
    if bucket.CountEvents() != 0 {
        tests.Fatal("Didn't get it")
    }

    bucket.AddEvent(event)
    if ! bucket.ContainsEvent(event) {
        tests.Fatal("Event not found")
    }

    if bucket.CountEvents() != 1 {
        tests.Fatal("Didn't get it")
    }

    bucket.RemoveEvent(event)
    if bucket.ContainsEvent(event) {
        tests.Fatal("Holy shit why is it still there")
    }

    event2 := &core.Event{
        Id: "345",
        Due: time.Now(),
        Payload: "Holy Shit",
        Expiry: time.Now().Add(time.Second * 60),
        EndPoint: "http://gc.com/fooooo",
    }

    bucket.AddEvent(event)
    bucket.AddEvent(event2)
    if bucket.CountEvents() != 2 {
        tests.Fatal("Didn't get it")
    }
}

func TestStoreEvent(tests *testing.T) {
    event := &core.Event{
        Id: "123",
        Due: time.Now(),
        Payload: "Holy Shit",
        Expiry: time.Now().Add(time.Second * 60),
        EndPoint: "http://gc.com/fooooo",
    }

    event2 := main.RetrieveEventById(event.Id)
    if event2 != nil {
        tests.Fatal("Oh snap what's up with the event being there?")
    }

    main.StoreEvent(event)
    event = main.RetrieveEventById(event.Id)
    if event == nil {
        tests.Fatal("Shit, didn't make it into the map")
    }

    bucket := main.FindBucketByTime(event.Due)
    if bucket == nil {
        tests.Fatal("The world is insane")
    }

    if ! bucket.ContainsEvent(event) {
        tests.Fatal("Wait, where is my event then?", bucket)
    }

    main.ClearEvent(event.Id)

    bucket = main.FindBucketByTime(event.Due)

    if bucket != nil {
        tests.Fatal("shoit")
    }
}
