.PHONY:gen format lint

gen:
	mockery

lint:
	golangci-lint run

format:
	golangci-lint fmt