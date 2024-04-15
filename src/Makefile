all: run test

.PHONY: test run up stop clean prepare_env

test: run
	go clean -testcache
	go test ./infrastructure/controllers -v

run: prepare_env build up

build:
	docker-compose build

up:
	docker-compose up --detach

stop:
	docker-compose stop

clean:
	docker-compose down

prepare_env:
	$(shell if ! [[ -f .env ]]; then cp env-example .env; fi;)
