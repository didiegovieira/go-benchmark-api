CMD_DIR = ./cmd

test:
	go test -coverprofile coverprofile.out -v ./... && go tool cover -html=coverprofile.out -o=coverprofile.html

wire:
	wire gen ./di

gen-mocks:
	./gen-mocks.sh

init:
	docker-compose up --build