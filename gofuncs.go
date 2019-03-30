package main

import (
	"C"
	"log"
	"time"
)

//export CanSleep
func CanSleep() C.int {
	time.Sleep(5 * time.Second)
	return 1
}

//export WillWake
func WillWake() {
	time.Sleep(5 * time.Second)
	log.Printf("Will Wake")
}

//export WillSleep
func WillSleep() {
	time.Sleep(5 * time.Second)
	log.Printf("Will Sleep")
}
