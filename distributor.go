package main

import "time"

func Tick() {
    for true {
        time.Sleep(ProbabilisticSleepDuration())
        Distribute(NextDueEvent())
    }
}

func Distribute(event Event) {
}
