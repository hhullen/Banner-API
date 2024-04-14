all: test run

test:
	go test ./infrastructure/controllers

run: prepare_env
	docker-compose build
	docker-compose up

prepare_env:
	$(shell if ! [[ -f .env ]]; then cp env-example .env; fi;)
