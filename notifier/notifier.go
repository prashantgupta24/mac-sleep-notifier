package main

import "log"

var instance *Notifier
var notifierCh chan *Activity

func GetInstance() *Notifier {
	if instance == nil {
		instance = &Notifier{}
	}
	return instance
}

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

func (n *Notifier) Quit() {
	n.quit <- struct{}{}
}
