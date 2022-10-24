.DEFAULT_GOAL = help
target=dev
tag=latest

## —— 📦 The amazing Makefile 📦 —————————————————————————————————————————

help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' Makefile | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— 🐋 Docker 🐋 —————————————————————————————————————————

up: ## Up the docker environment
	docker network ls | grep pfm-network > /dev/null || docker network create pfm-network
	docker-compose -f dev-docker-compose.yml up -d
down: ## Down the docker environment
	docker-compose -f dev-docker-compose.yml down

## —— database —————————————————————————————————————————

confirm:
	@( read -p "$(RED)Are you sure? [y/N]$(RESET): " sure && case "$$sure" in [yY]) true;; *) false;; esac )

