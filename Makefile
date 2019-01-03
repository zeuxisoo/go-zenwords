all:
	@echo
	@echo "Commands   : Description"
	@echo "---------- : -----------------"
	@echo "make tools : Fetch related development tools"
	@echo

tools:
	go get -u github.com/golang/protobuf/protoc-gen-go

generate:
	protoc --proto_path=./rpc/protos/ --go_out=plugins=grpc:rpc/protos ./rpc/protos/*.proto
