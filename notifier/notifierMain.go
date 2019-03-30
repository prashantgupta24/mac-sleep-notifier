package main

// #cgo LDFLAGS: -framework CoreFoundation -framework IOKit
// int CanSleep();
// void WillWake();
// void WillSleep();
// #include "main.h"
import "C"
import "fmt"

func StartNotifier() {
	C.registerNotifications()
	fmt.Println("stopping notifier")
}

func StopNotifier() {
	C.unregisterNotifications()
}
