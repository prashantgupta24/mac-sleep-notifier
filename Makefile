all: build

build: clean
	go build -o ./bin/sleep-notifier ./notifier/*.go && ./bin/sleep-notifier

clean:
	rm -rf ./bin