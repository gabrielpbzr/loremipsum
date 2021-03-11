build:
	go build
clean:
	rm -f loremipsum
test: build
	go test -v ./...
install: test
	cp loremipsum /usr/local/bin/
