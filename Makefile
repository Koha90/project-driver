run:
	@./bin/project-driver

build:
	@go build -o bin/

test:
	test -v ./...
