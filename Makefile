.PHONY: test run all
test:
	go test -v .

run:
	go run test/main.go

all: test run