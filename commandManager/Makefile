protoGen:
	protoc 	app/Adapter/In/ApiGrcp/proto/*.proto --go_out=app/Adapter/In/ApiGrcp/gen --go-grpc_out=app/Adapter/In/ApiGrcp/gen
clearGit:
	git branch --merged | egrep -v "(^\*|master|main|dev)" | xargs git branch -d
tests:
	go test ./...