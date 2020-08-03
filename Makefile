# Path to .proto files
PROTO_PATH := carpb

# Output directory
GRPC_OUT := generated

protoc:
	# Generate proto stubs
	protoc \
	--go_out=plugins=grpc:. \
	--go_opt=paths=source_relative \
	$(PROTO_PATH)/*.proto