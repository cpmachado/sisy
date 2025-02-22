build:
	go build ./cmd/sisy

clean:
	rm -f sisy

lint:
	golangci-lint run .

setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5
	go install golang.org/x/vuln/cmd/govulncheck@latest

.PHONY: build clean lint setup
