default: docker-test

# Local commands
.PHONY: test
test:
	go test

.PHONY: coverage
coverage:
	@go test -coverprofile=/tmp/coverage.out
	@go tool cover -html=/tmp/coverage.out -o ./coverage.html
	echo "Open coverage.html"

.PHONY: mutation
mutation:
	@go-mutesting .

.PHONY: docker-test
docker-test:
	@docker run --rm -v ${PWD}:/code -w /code golang:1.18 make test

.PHONY: docker-coverage
docker-coverage:
	@docker run --rm -v ${PWD}:/code -w /code golang:1.18 make coverage

.PHONY: docker-mutation
docker-mutation:
	@docker run --rm -v ${PWD}:/code codiumteam/legacy-training-go make mutation
