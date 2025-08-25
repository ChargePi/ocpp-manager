.PHONY: gen format lint test

test:
	go test -v -race -covermode=atomic ./... -coverpkg=./... -short -coverprofile=unit_coverage.out

gen:
	mockery

lint:
	golangci-lint run

format:
	golangci-lint fmt