protoGen:
	protoc 	app/Adapter/In/Grcp/proto/*.proto --go_out=app/Adapter/In/Grcp/gen --go-grpc_out=app/Adapter/In/Grcp/gen
clearGit:
	git branch --merged | egrep -v "(^\*|master|main|dev)" | xargs git branch -d