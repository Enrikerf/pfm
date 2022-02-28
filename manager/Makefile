args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
#COLORS
GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RED := $(shell tput -Txterm setaf 1)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: help dependencies up start stop restart status ps clean log console flush sync aws reset

.DEFAULT_GOAL := help

protoGen: ##@proto Start all or <container_name> containers in foreground
	protoc 	app/Adapter/In/ApiGrcp/proto/*.proto --go_out=app/Adapter/In/ApiGrcp/gen --go-grpc_out=app/Adapter/In/ApiGrcp/gen

clearGit: ##@git Start all or <container_name> containers in foreground
	git branch --merged | egrep -v "(^\*|master|main|dev)" | xargs git branch -d

runTests: ##@tests Start all or <container_name> containers in foreground
	cd tests; go test ./... -v; cd -

migrate: ##@migrate Start all or <container_name> containers in foreground
	cd app/Adapter/Out/Persistence/Migrations/ ; migrate create -ext sql -dir . -seq $(call args)

migrate_up:
	migrate -path app/Adapter/Out/Persistence/Migrations/ -database "mysql://root:poiu@tcp(localhost:3306)/go_project" -verbose up

migrate_down:
	migrate -path app/Adapter/Out/Persistence/Migrations/ -database "mysql://root:poiu@tcp(localhost:3306)/go_project" -verbose down

help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

evans:
	./evans.sh -p 9090 -r

HELP_FUN = \
    %help; \
    while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
    print "usage: make [target]\n\n"; \
    for (sort keys %help) { \
    print "${WHITE}$$_:${RESET}\n"; \
    for (@{$$help{$$_}}) { \
    $$sep = " " x (32 - length $$_->[0]); \
    print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
    }; \
    print "\n"; }