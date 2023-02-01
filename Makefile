build:
	@go build -o ./bin/blockchain-project
run: build
	@./bin/blockchain-project
test:
	go test -v ./...