.PHONY: all benchmark certificate clean details html test

all: clean test

benchmark: clean
	go test -benchtime 200ms -benchmem -bench . benchmark_test.go

certificate:
	openssl req -new -newkey "rsa:2048" -days 365 -nodes -x509 -keyout "server.key" -out "server.crt"

clean:
	git clean -fX

details: coverage.test
	go tool cover -func=$<

html: coverage.test
	go tool cover -html=$<

test:
	go test -cover -count=1 ./...

%.txt %.test:
	go test -cover -count=1 -coverprofile=$@ ./...
