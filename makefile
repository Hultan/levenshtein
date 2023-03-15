all: build test

build:
#    go build -o build/levenshtein levenshtein.go

run: build
#    ./build/levenshtein

test:
#	go test -v ./... -count=1
	go test -v ./...

bench:
	go test -bench=.