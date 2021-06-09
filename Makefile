GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BINARY_NAME=hciengserver
MAIN=src/cmd/hciengserver/main.go

build:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) $(MAIN)

test:
	$(GOTEST) ./... -coverprofile coverage.txt
	go tool cover -func coverage.txt

clean:
	$(GOCLEAN)
	rm -f ./bin/$(BINARY_NAME)

run:
	$(GORUN) $(MAIN)

deps:
	$(GOGET) github.com/hcieng/hcieng-server
	hcieng-server sync

fmt:
	go fmt ./src/...