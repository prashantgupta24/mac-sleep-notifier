package main

import (
	"log"

	"github.com/prashantgupta24/mac-sleep-notifier/notifier"
)

func main() {
	log.Printf("starting sleep notifier ...")
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
}
