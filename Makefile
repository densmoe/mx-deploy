BINARY_NAME=mx-deploy

build:
	CGO_ENABLED=0 go build -o ./mx-deploy

install: build
	cp $(BINARY_NAME) /usr/local/bin/