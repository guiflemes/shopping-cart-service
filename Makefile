include .env
export

.PHONY: proto
proto:
	@./scripts/proto.sh shoppingCartService