[![Go Report Card](https://goreportcard.com/badge/github.com/prashantgupta24/mac-sleep-notifier)](https://goreportcard.com/report/github.com/prashantgupta24/mac-sleep-notifier) [![codecov](https://codecov.io/gh/prashantgupta24/mac-sleep-notifier/branch/master/graph/badge.svg)](https://codecov.io/gh/prashantgupta24/mac-sleep-notifier) [![version][version-badge]][RELEASES] [![godoc-badge][godoc-badge]][godoc-link]

## macOS Sleep/ Wake notifications in golang

Inspired from [this](https://nicolai86.eu/blog/2017/12/sleep-wake-notifications-in-go/) blog.

This libary notifies through a channel whenever your machine is put to sleep or is woken up. Calling the `Start` function will get you a channel on which you receive both `Sleep` and `Awake` activities.

## Installation

`go get -u github.com/prashantgupta24/mac-sleep-notifier/notifier`

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
## Example
See example [here](https://github.com/prashantgupta24/mac-sleep-notifier/blob/master/example/example.go)

[godoc-badge]: https://img.shields.io/badge/godoc-reference-blue.svg
[godoc-link]: https://godoc.org/github.com/prashantgupta24/mac-sleep-notifier/notifier

[version-badge]: https://img.shields.io/github/release/prashantgupta24/mac-sleep-notifier.svg
[RELEASES]: https://github.com/prashantgupta24/mac-sleep-notifier/releases
