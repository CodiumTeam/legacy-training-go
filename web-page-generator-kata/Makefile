default: docker-run

.PHONY: run
run:
	go run main.go
	@printf "Open html/users.html\n"

.PHONY: docker-run
docker-run:
	@docker run --rm -v ${PWD}:/code -w /code golang:1.18 make run
