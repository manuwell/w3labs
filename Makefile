install:
	@go install github.com/vektra/mockery/v3
	@go mod download

mock-gen: install
	@mockery 

docker-up:
	@docker-compose up -d

server: install docker-up
	cd ./bin && sleep 5 && ./run.sh

