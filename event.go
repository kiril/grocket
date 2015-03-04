package main

import (
	"time"
)


type Event struct {
	Id          string    `json:"id"`
	Due         time.Time `json:"due"`
	Payload     string    `json:"payload"`
	Expiry      time.Time `json:"expiry"`
	EndPoint    string    `json:"endpoint"`
	MaxAttempts int       `json:"max-attempts"`
	Verb        string    `json:"verb"`
}

type Events []Event
