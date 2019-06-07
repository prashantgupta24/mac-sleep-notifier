COVER_PROFILE=cover.out
COVER_HTML=cover.html

.PHONY: $(COVER_PROFILE) $(COVER_HTML)

all: coverage vet

coverage: $(COVER_HTML)

$(COVER_HTML): $(COVER_PROFILE)
	go tool cover -html=$(COVER_PROFILE) -o $(COVER_HTML)

$(COVER_PROFILE):
	go test -v -failfast -race -coverprofile=$(COVER_PROFILE) ./...

vet:
	go vet $(shell glide nv)

test:
	go test -v -failfast -race ./...

example-start:
	go run example/example.go