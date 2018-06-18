NAME=sigil
ARCH=$(shell uname -m)
VERSION=0.5.0

PLATFORMS := darwin linux
os = $(word 1, $@)

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p build/$(os)
	go build -a -o build/$(os)/sigil ./cmd

test:
	cd builtin ; go test ; cd ..
	basht tests/*.bash

install: build
	install build/$(shell uname -s)/sigil /usr/local/bin

deps:
	go get -u github.com/progrium/basht/...
	go get -d ./cmd

clean:
	rm -rf build release

.PHONY: build release
