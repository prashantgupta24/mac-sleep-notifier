/*
Package notifier notifies through a channel whenever your machine is put to sleep or is woken up. Calling
the Start() function will get you a channel on which you receive both "Sleep" and "Awake" activities.

Installation

The library can be installed using:

 go get -u github.com/prashantgupta24/mac-sleep-notifier/notifier

Usage

The usage is as following:

 notifierCh := notifier.GetInstance().Start()

 for {
	select {
	case activity := <-notifierCh:
		if activity.Type == notifier.Awake {
			log.Println("machine awake")
		} else {
			if activity.Type == notifier.Sleep {
				log.Println("machine sleeping")
			}
		}
	}
 }

*/
package notifier
