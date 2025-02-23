default: build

.PHONY: build
build:
	@echo
	go build ./main.go
	@echo

.PHONY: run
run:
	@clear
	@go run ./main.go
	@echo

.PHONY: clean
clean:
	@echo
	go clean
	@echo

.PHONY: test
test:
	@echo
	go test
	@echo

.PHONY: fmt
fmt:
	@echo
	go fmt ./...
	@echo
