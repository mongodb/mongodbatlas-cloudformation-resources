.PHONY: build test clean

build:
	cfn generate || cat rpdk.log
	env GOOS=linux go build -ldflags="-s -w" -tags="logging callback scheduler" -o bin/handler cmd/main.go

test:
	cfn generate
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler cmd/main.go

clean:
	rm -rf bin
