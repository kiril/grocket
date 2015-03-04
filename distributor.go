package main

import "time"

func ProcessBacklog() {
    // this whole thing has to happen transactionally...
    // which I don't even know how to do yet.
    for bucket := NextTimeBucket(); bucket != nil && time.Now().After(bucket.Time); {
        for i := 0; i < len(bucket.EventIds); i++ {
            event := RetrieveEventById(string(bucket.EventIds[i]))
            Distribute(event)
        }
        RemoveTimeBucket(bucket)
    }

    // what I need to do is get this to happen again in a new block? or something.
    // time.Sleep(ProbabilisticSleepDuration())
}

func Distribute(event *Event) {

}
