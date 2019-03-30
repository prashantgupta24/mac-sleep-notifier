package main

// #cgo LDFLAGS: -framework CoreFoundation -framework IOKit
// int CanSleep();
// void WillWake();
// void WillSleep();
// #include "main.h"
import "C"
import (
	"log"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		log.Printf("waiting ...")
		C.registerNotifications()
		log.Printf("Go bye")
		wg.Done()
	}(&wg)

	log.Printf("Hello, World")
	wg.Wait()
	// C.unregisterNotifications()
	// time.Sleep(time.Minute)
}
