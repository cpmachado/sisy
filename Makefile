SRC=$(wildcard **/*.go)
BIN=sisy

build: ${BIN}
	@echo sisy built

${BIN}: ${SRC}
	go build -o ${BIN} ./cmd/sisy

clean:
	rm -f ${BIN} cover.out

lint:
	golangci-lint run ./...

test:
	go test -v -cover -coverprofile=cover.out ./...

.PHONY: build clean lint test
