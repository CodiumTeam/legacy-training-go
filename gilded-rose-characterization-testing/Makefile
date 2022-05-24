default: docker-test

# Local commands
.PHONY: test
test:
	go test

.PHONY: coverage
coverage:
	@go test -coverprofile=/tmp/coverage.out
	@go tool cover -html=/tmp/coverage.out -o /code/coverage.html
	echo "Open coverage.html"

.PHONY: docker-test
docker-test:
	@docker run --rm -v ${PWD}:/code -w /code golang:1.18 make test

.PHONY: docker-coverage
docker-coverage:
	@docker run --rm -v ${PWD}:/code -w /code golang:1.18 make coverage
