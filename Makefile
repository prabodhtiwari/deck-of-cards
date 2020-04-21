.PHONY: test

run:
	go run cmd/server.go

unit-tests:
	(cd test/unit && go test -v);

integration-tests:
	(cd test/integration && go test -v);

test:	unit-tests integration-tests

