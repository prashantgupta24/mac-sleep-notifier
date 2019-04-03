## macOS Sleep/ Wake notifications in golang

Inspired from [this](https://nicolai86.eu/blog/2017/12/sleep-wake-notifications-in-go/) blog.

This libary notifies through a channel whenever your machine is put to sleep or is woken up. Calling the `Start` function will get you a channel on which you receive both `Sleep` and `Awake` activities.

## Installation

`go get -u github.com/prashantgupta24/mac-sleep-notifier`

## Usage

```go
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
```