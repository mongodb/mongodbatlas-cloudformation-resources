.PHONY: build test clean

build:
	make -f makebuild  # this runs build steps required by the cfn cli

test:
	cfn generate
	env GOOS=linux go build -ldflags="-s -w" -tags="lambda.norpc,$(TAGS)" -o bin/bootstrap cmd/main.go

clean:
	rm -rf bin
