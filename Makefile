PKG=github.com/Code-Hex/grpc-gate
OUTPUT_DIR=_output

proto/compile:
	mkdir -p $(OUTPUT_DIR)
	protoc -I. --go_out=plugins=grpc:$(OUTPUT_DIR) internal/proto/*.proto
	cp $(OUTPUT_DIR)/$(PKG)/internal/proto/*.go internal/proto/
	rm -rf $(OUTPUT_DIR)

proto/clean:
	rm -f internal/proto/*.pb.go
