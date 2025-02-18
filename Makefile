build_proto:
	protoc --go_out=./notificationservice/notificationproto --go_opt=paths=source_relative \
    --go-grpc_out=./notificationservice/notificationproto --go-grpc_opt=paths=source_relative \
    ./notificationservice/notificationproto/service.proto