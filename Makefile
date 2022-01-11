.EXPORT_ALL_VARIABLES:

PORT=8080
SERVER_CLIENT=stage.nekom.com:443
SERVER_TOKEN=storer@2020!
SERVER_CLIENT_UUID=fdfd497828effdfda2b049a27849a2d4

gen:
	protoc --go_out=./pb --go-grpc_out=./pb proto/*

run:
	go run cmd/server/server.go