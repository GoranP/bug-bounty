.DEFAULT_GOAL := default

clean:
	@rm ./bin/*

default: bin_dir build

build: build
	@go get
	@go build -o bin/innsecure ./cmd/innsecure
	@go build -o bin/token ./cmd/token

bin_dir:
	@mkdir -p ./bin
