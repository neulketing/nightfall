.PHONY: build test clean install

BINARY = nightfall
GO = go

build:
	$(GO) build -o $(BINARY) ./cmd/nightfall

test:
	$(GO) test ./cmd/... ./core/... ./recon/... ./initial-access/... ./c2/... ./ai/... ./integration/... -v

clean:
	rm -f $(BINARY)
	rm -rf bin/

install: build
	install -m 755 $(BINARY) /usr/local/bin/$(BINARY)
