#!/bin/bash

build-production:
	@echo "Building for production..."
	@go build
	@./nyanyangku -env=production

build-development:
	@echo "Building with debug..."
	@go build
	@./nyanyangku

test:
	@echo "Unit Testing..."
	@go test -cover ./...
