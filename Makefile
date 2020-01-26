IMAGE_NAME ?= "accountapi-client"
DOCKER_COMPOSE = docker-compose -f docker-compose.yml
CWD = $(shell pwd)

.PHONY: test unit-tests bdd-tests clean stop start help

help:			## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

start:			## Brings the test environment up and runs all the tests.
	- make stop
	${DOCKER_COMPOSE} up

start-services:		## Brings the test environment up for developers running locally.
	${DOCKER_COMPOSE} up -d vault postgresql accountapi
	- docker ps -a

stop:			## Stops any running containers for the test environment.
	${DOCKER_COMPOSE} down -v
	- docker ps -a

clean:			## Cleans any unwanted stuff
	go clean
	go clean -testcache

bdd-tests:		## Runs the BDD tests locally
	@echo
	@echo "---------------------------------------- Running BDD Tests"
	@echo
	curl -f -s "${URL_UNDER_TEST}/v1/health"
	go get github.com/DATA-DOG/godog/cmd/godog
	cd test ; godog *.feature
	@echo
	@echo "!!! ALL TESTS GOOD !!!"

unit-tests:		## Runs unit tests locally
	@echo
	@echo "---------------------------------------- Running Unit Tests"
	@echo
	go vet ./...
	go clean -testcache
	go test -v --cover ./accountclient/...

local-test: 		## Run all tests - command for docker-compose or run `make start-services` first
	./scripts/wait-for-api.sh
	make unit-tests
	make bdd-tests

run-example:		## Runs the example reference implementation.
	- rm ./example
	go build ./example.go
	./example
