default: docker-run

.PHONY: run
run:
	go run main.go
	@printf "Open html/users.html\n"

.PHONY: test
test:
	go test

.PHONY: docker-run
docker-run:
	@docker run --rm -v ${PWD}:/code -w /code golang:1.18 make run

.PHONY: docker-test
docker-test:
	@docker run --rm -v ${PWD}:/code -w /code golang:1.18 make test
