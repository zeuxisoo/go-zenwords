all:
	@echo
	@echo "Commands   : Description"
	@echo "---------- : -----------------"
	@echo "make tools : Fetch related development tools"
	@echo

tools:
	go get -u github.com/smartystreets/goconvey
	go get -u github.com/golang/protobuf/protoc-gen-go

generate:
	protoc --proto_path=./rpc/proto/ --go_out=plugins=grpc:rpc/proto ./rpc/proto/*.proto
