.PHONY: gen-go gen-notification-go gen-users-go

gen-go: gen-users-go gen-notification-go

gen-users-go:
	protoc --go_out=./go/users --go_opt paths=source_relative \
		   --go-grpc_out=./go/users --go-grpc_opt paths=source_relative \
		   --proto_path=./proto users.proto

gen-notification-go:
	protoc --go_out=./go/notification --go_opt paths=source_relative \
		   --go-grpc_out=./go/notification --go-grpc_opt paths=source_relative \
		   --proto_path=./proto notification.proto

