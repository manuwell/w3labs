install:
	@go install github.com/vektra/mockery/v3
	@go mod download

mock-gen: install
	@mockery 

server: install 
	cd ./bin && ./run.sh

