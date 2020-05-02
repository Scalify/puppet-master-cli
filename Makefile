.PHONY: build
build:
	@docker build -t puppet-master-cli .

.PHONY: vendors
vendors:
	go mod download
	go mod tidy

.PHONY: test
test:
	go test -cover ./...
	golangci-lint run
	golint -set_exit_status ./...

.PHONY: run-example
run-example:
	go run main.go exec --executor-logs-verbose \
		--code example/code.mjs \
		--module example/modules/shared.mjs \
		--vars example/vars.json
