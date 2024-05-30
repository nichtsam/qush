.PHONY: start build

start: build
	@./build/qush

build:
	@go build -o ./build/qush .
