#!make

.PHONY: default local-serve install

default: local-serve

install:
	cp .docker/golang/.env.example .docker/golang/.env
	docker-compose build
	dep ensure

local-serve:
	GRPC_ADDRESS=:50051 go run main.go