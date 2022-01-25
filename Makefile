.PHONY: build
build:
	go build -o bin/cleaner cmd/cleaner/main.go

.PHONY: all
all:
	@go build -o bin/cleaner cmd/cleaner/main.go
	@echo 'The app was successfully built at ./bin/cleaner '
	@./bin/cleaner -h

.PHONY: run
run:
	go run main.go

.PHONY: fake
fake:
	go run cmd/faker/main.go

.PHONY: test
test:
	go test -covermode=atomic ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	@echo '>> cleaning go'
	@go clean
	@echo '>> cleaning binaries'
	@-rm -rf bin
	@echo '>> cleaning fake files'
	@-rm -rf files
