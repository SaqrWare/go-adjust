all: test build

test:
	go test -v
build:
	go build -o bin/myhttp
watch:
	ag -l |grep .go |entr -s 'make build'
clean:
	rm -r bin/myhttp