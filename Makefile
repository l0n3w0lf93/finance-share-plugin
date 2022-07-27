PROTO_FILES=$(shell find proto -name *.proto)
.PHONY: plugin
# generate plugin proto
protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       $(PROTO_FILES)
