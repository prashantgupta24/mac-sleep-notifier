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

	go func() { StartNotifier() }()

	go func(n *Notifier) {
		for {
			select {
			case <-n.quit:
				log.Printf("quitting notifier")
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
