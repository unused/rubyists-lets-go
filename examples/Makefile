CMD=cmd/main.go
BIN=bin/gorgonzola

.PHONY: build
build:
	@go build -o $(BIN) $(CMD)
	@./$(BIN)

.PHONY: run
run:
	@go run $(CMD)

.PHONY: test
test:
	@go test

.PHONY: counter
counter:
	go run cmd/count-server.go

.PHONY: clean
clean:
	rm -f $(BIN)
