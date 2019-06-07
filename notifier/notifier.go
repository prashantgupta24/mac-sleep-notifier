package notifier

import "log"

var instance *Notifier
var notifierCh chan *Activity

//GetInstance gets the singleton instance for the notifier
func GetInstance() *Notifier {
	if instance == nil {
		instance = &Notifier{}
	}
	return instance
}

//Start the notifier. It returns an Activity channel
//to listen for machine sleep/wake activities.
func (n *Notifier) Start() chan *Activity {

	n.quit = make(chan struct{})
	notifierCh = make(chan *Activity)

	go func(n *Notifier) {
		n.setIsRunning(true)
		StartNotifier()
	}(n)

	go func(n *Notifier) {
		for {
			select {
			case <-n.quit:
				log.Printf("quitting notifier")
				n.setIsRunning(false)
				StopNotifier()
				return
			}
		}
	}(n)
	return notifierCh
}

//Quit the notifier
func (n *Notifier) Quit() {
	n.quit <- struct{}{}
}

//setIsRunning sets status of notifier
func (n *Notifier) setIsRunning(status bool) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.isRunning = status
}

//isStatusRunning checks running status of notifier
func (n *Notifier) isStatusRunning() bool {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	return n.isRunning
}
