package notifier

import "sync"

//Notifier notifies about the sleep/wake events
type Notifier struct {
	quit      chan struct{}
	mutex     sync.RWMutex
	isRunning bool
}

//Type determines if it is a sleep or an awake type activity
type Type string

/*
The 2 types of activities
*/
const (
	Sleep Type = "sleep"
	Awake Type = "awake"
)

//Activity struct is used to hold the sleep and awake activity.
type Activity struct {
	Type Type
}
