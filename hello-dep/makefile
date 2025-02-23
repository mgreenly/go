
default: checks test

.PHONY: test
test:
	@echo
	go test ./...

.PHONY: checks
checks:
	@echo
	staticcheck ./...

.PHONY: fmt
fmt:
	@echo
	go fmt ./...
	@echo
