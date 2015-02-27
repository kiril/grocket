package main

import "time"

func Tick() {
    time.Sleep(ProbabilisticSleepDuration())
    event := NextDueEvent()
    Distribute(&event)
}

func Distribute(event *Event) {

}
