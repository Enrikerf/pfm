.DEFAULT_GOAL := help
target=dev
tag=latest

## —— 📦 Makefile 📦 —————————————————————————————————————————

help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— 🐋 Docker 🐋 —————————————————————————————————————————

build: ## Up the docker environment
	target=${target} docker-compose build --progress=plain
up: ## Up the docker environment
	target=${target} docker-compose up -d
down: ## Down the docker environment
	target=${target} docker-compose down
release:
	target=prod docker-compose build --no-cache --progress=plain
run-prod:
	docker run --rm --name app-prod go-dev:19.2-prod

## —— Go Container —————————————————————————————————————————
enter: ## Access to container
	docker exec -u docker-user -ti go-manager /usr/bin/fish
run: ## run main.go
	docker exec -ti go-dev go run main.go
debug: ## run on debug mode, remember to listen on the IDE
	docker exec -ti go-manager /go/bin/dlv --headless --listen=:40000 --api-version=2 debug
runTests: ##@tests Start all or <container_name> containers in foreground
	cd tests; go test ./... -v; cd -

## —— Proto —————————————————————————————————————————
protoGen: ## @proto Start all or <container_name> containers in foreground
	protoc 	app/Adapter/In/ApiGrcp/proto/*.proto --go_out=app/Adapter/In/ApiGrcp/gen --go-grpc_out=app/Adapter/In/ApiGrcp/gen

## —— Git —————————————————————————————————————————
clearGit: ## @git Start all or <container_name> containers in foreground
	git branch --merged | egrep -v "(^\*|master|main|dev)" | xargs git branch -d



## —— Database —————————————————————————————————————————

migrate: ## migrate
	cd app/Adapter/Out/Persistence/Migrations/ ; migrate create -ext sql -dir . -seq $(call args)

migrate_up: ## migrate_up
	migrate -path app/Adapter/Out/Persistence/Migrations/ -database "mysql://root:poiu@tcp(localhost:3306)/go_project" -verbose up

migrate_down: ## migrate_down
	migrate -path app/Adapter/Out/Persistence/Migrations/ -database "mysql://root:poiu@tcp(localhost:3306)/go_project" -verbose down

## —— Evans —————————————————————————————————————————
evans: ## evans
	./evans.sh -p 9090 -r

