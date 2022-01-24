build:
	go build -o bin/cleaner cmd/cleaner/main.go

all:
	go build -o bin/cleaner cmd/cleaner/main.go
	./bin/app

run:
	go run main.go

test:
	go test -covermode=atomic ./...

lint:
	golangci-lint run

clean:
	go clean
	rm bin/app
