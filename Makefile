.PHONY: build test clean install

BINARY = nightfall
GO = go

build:
	$(GO) build -o $(BINARY) ./cmd/nightfall

test:
	$(GO) test ./... -v

clean:
	rm -f $(BINARY)
	rm -rf bin/

install: build
	install -m 755 $(BINARY) /usr/local/bin/$(BINARY)
