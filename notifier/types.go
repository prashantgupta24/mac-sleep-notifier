package main

//Notifier notifies about the sleep/wake events
type Notifier struct {
	quit chan struct{}
}

type Type string

const (
	Sleep Type = "sleep"
	Awake Type = "awake"
)

type Activity struct {
	Type Type
}
