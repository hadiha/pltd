.PHONY : build run

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

proto:
	$(foreach proto_file, $(shell find api/proto -name '*.proto'),\
	protoc --proto_path=api/proto --go_out=plugins=grpc:api/proto \
	--go_opt=paths=source_relative $(proto_file);)

migration:
	@go run cmd/migration/migration.go $(call args,up)

build:
	@go build -o bin

run: build
	@./bin

docker:
	docker build -t PLTD3:latest .

run-container:
	docker run --name=PLTD3 --network="host" -d PLTD3

clear-docker:
	docker rm -f PLTD3
	docker rmi -f PLTD3

# mocks all interfaces for unit test
mocks:
	@mockery --all --keeptree --dir=internal --output=pkg/mocks --case underscore
	@mockery --all --keeptree --dir=pkg --output=pkg/mocks --case underscore

# unit test & calculate code coverage
test:
	@echo "\x1b[32;1m>>> running unit test and calculate coverage\x1b[0m"
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count \
		-coverpkg=$$(go list ./... | grep -v mocks | tr '\n' ',')
	@go tool cover -func=coverage.txt
