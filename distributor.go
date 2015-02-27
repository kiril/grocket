package main

import "time"

func Tick() {
    time.Sleep(ProbabilisticSleepDuration())
    Distribute(NextDueEvent())
}

func Distribute(event Event) {
}
