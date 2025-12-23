PROTO_DIR = grpc
PROTO_FILE = server.proto
PROTO_PATH = $(PROTO_DIR)/$(PROTO_FILE)

OUT_DIR = grpc


# Compile protobuf files
proto: $(OUT_DIR) 
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative -I=$(PROTO_DIR) $(PROTO_FILE)

clean:
	rm -rf $(OUT_DIR)/*.pb.go

datafeed:
	go build cmd/datafeed/main.go -o bin/datafeed
	scp -C  bin/datafeed root@51.79.177.97:/root/datafeed/