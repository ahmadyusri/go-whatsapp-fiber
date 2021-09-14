#!make
include .env
export $(shell sed 's/=.*//' .env)

.PHONY: test security run stop

BUILD_DIR = $(PWD)/app

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

swag:
	swag init

docker_build_image:
	docker-compose up -d --build

docker_run_image:
	docker-compose up -d

docker_app: docker_build_image

run: swag docker_app

stop:
	docker-compose down
