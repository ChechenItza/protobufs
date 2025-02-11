gen_grpc:
	protoc \
	 --go_out=booking \
	 --go_opt=paths=source_relative \
	 --go-grpc_out=booking \
	 --go-grpc_opt=paths=source_relative \
	 booking.proto
