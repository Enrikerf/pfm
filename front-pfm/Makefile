args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
#COLORS
GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RED := $(shell tput -Txterm setaf 1)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: help dependencies up start stop restart status ps clean log console flush sync aws reset

.DEFAULT_GOAL := help
#
#protoGen: ##@proto Start all or <container_name> containers in foreground
#	protoc --proto_path=src/protobuf \
#		--js_out=import_style=commonjs,binary:src/protobuf/generated \
#		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src/protobuf/generated \
#	src/protobuf/*.proto

protoGen: ##@proto Start all or <container_name> containers in foreground
	protoc --proto_path=src/protobuf --js_out=import_style=commonjs,binary:src/protobuf/generated --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src/protobuf/generated src/protobuf/*.proto
	#protoc --proto_path=src/protobuf --grpc-web_out=import_style=typescript,mode=grpcwebtext:./src/protobuf/generated src/protobuf/*.proto
help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

evans:
	./evans.sh -p 8081 -r

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