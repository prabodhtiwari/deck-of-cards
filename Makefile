.PHONY: test

run:
	go run cmd/server.go

unit-tests:
	go test -v github.com/deck-of-cards/utils

integration-tests:
	go test -v github.com/deck-of-cards/api/handlers

test:	unit-tests integration-tests

