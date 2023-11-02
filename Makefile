BINARY_NAME="wisenotes"

build: clean generate
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) *.go

generate:
	@echo "Generating..."
	@go generate ./...

clean:
	@echo "Cleaning up..."
	@go mod tidy -v
	@rm -rf bin/
	@mkdir -p bin/

