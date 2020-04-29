PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

install:
	go mod download

build:
	go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/*.go || exit
	
tests:
	go test ./cmd/$(PROJECTNAME)/*.go

docker-build:
	docker-compose -f ./deploy/docker-compose.yaml build remotelogs-api

docker-run: docker-build
	docker-compose -f ./deploy/docker-compose.yaml up remotelogs-api

production:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go

start:
	go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/*.go || exit
	./bin/$(PROJECTNAME)
