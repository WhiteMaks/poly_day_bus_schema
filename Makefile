up-test-environment:
	docker-compose up -d
	sleep 60

test:
	go clean -testcache && go test -cover -v ./...

build:
	go build -o ./.bin/bus-schema/poly_day_bus_schema cmd/main.go

.PHONY: build test up-test-environment