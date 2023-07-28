.PHONY: build-local

build-local: export GOARCH=amd64
build-local: plugins/vault-aelf

plugins/vault-aelf: cmd/vault-plugin-secrets/main.go
	CGO_ENABLED=1 go build -ldflags="-s -w" -o $@ $<

clean:
	rm -f plugins/vault-aelf