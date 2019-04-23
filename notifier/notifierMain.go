package notifier

// #cgo LDFLAGS: -framework CoreFoundation -framework IOKit
// int CanSleep();
// void WillWake();
// void WillSleep();
// #include "main.h"
import "C"

//StartNotifier starts the internal notifier function which communicates with the C library.
func StartNotifier() {
	C.registerNotifications()
}

//StopNotifier stops the internal notifier function which communicates with the C library.
func StopNotifier() {
	C.unregisterNotifications()
}
