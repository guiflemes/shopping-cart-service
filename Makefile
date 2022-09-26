include .env
export

.PHONY: proto
proto:
	@./scripts/proto.sh shoppingCartService


all:down build up test

build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down --remove-orphans