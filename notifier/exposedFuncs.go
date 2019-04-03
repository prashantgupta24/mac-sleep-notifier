package notifier

import (
	"C"
	"time"
)

//export CanSleep
func CanSleep() C.int {
	time.Sleep(5 * time.Second)
	return 1
}

//export WillWake
func WillWake() {
	//time.Sleep(5 * time.Second)
	notifierCh <- &Activity{
		Type: Awake,
	}
	//log.Printf("Will Wake")
}

//export WillSleep
func WillSleep() {
	//time.Sleep(5 * time.Second)
	notifierCh <- &Activity{
		Type: Sleep,
	}
	//log.Printf("Will Sleep")
}
