docbuild:
	swag init --parseDependency --parseInternal -d ./,./pkg/common/models,./pkg/products,./pkg/users
	docker compose -f ./docker-compose.yml build

build:
	docker compose -f ./docker-compose.yml build

run:
	docker compose -f ./docker-compose.yml up

stop:
	docker compose -f ./docker-compose.yml down