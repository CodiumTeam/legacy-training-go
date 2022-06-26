default: docker-run

# Local commands
.PHONY: run
run:
	go run main.go

.PHONY: docker-run
docker-run:
	@docker run --rm -v ${PWD}:/code -w /code golang:1.18 make run

