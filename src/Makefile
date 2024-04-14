all: test run

test:
	go clean -testcache
	go test ./infrastructure/controllers -v

run: prepare_env
	docker-compose build
	docker-compose up

prepare_env:
	$(shell if ! [[ -f .env ]]; then cp env-example .env; fi;)
