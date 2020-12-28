clean:
	rm -r bin

test:
	go test ./... -coverprofile fmtcoverage.html fmt

build:
	mkdir -p bin
	go build -o bin/service