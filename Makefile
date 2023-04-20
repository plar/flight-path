SVC_NAME := flight-path

.PHONY: clean build test coverage

test:
	go test -v ./...

build:
	go build -o $(SVC_NAME)

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

clean:
	go clean -modcache
	rm -f $(SVC_NAME)
	rm -f coverage.out
	rm -f coverage.html

docker-build:
	docker build -t flight-path-svc .

# Run the service in a Docker container
docker-run:
	docker run --rm -p 8080:8080 flight-path-svc

curl-test:
	curl -X POST -H "Content-Type: application/json" \
		-d '[["SFO", "EWR"]]' \
		http://localhost:8080/calculate
	curl -X POST -H "Content-Type: application/json" \
		-d '[["ATL", "EWR"], ["SFO", "ATL"]]' \
		http://localhost:8080/calculate
	curl -X POST -H "Content-Type: application/json" \
		-d '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]' \
		http://localhost:8080/calculate
