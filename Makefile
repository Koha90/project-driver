build:
	@go build -o bin/

run: build
	@./bin/project-driver

test:
	test -v ./...
