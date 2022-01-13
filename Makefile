.EXPORT_ALL_VARIABLES:

PORT=8001
SERVER_CLIENT=stage.nekom.com:443
SERVER_TOKEN=storer@2020!
SERVER_CLIENT_UUID=fdfd497828effdfda2b049a27849a2d4

gen:
	protoc --go_out=./pb --go-grpc_out=./pb proto/*

run:
	go run cmd/server/server.go


docker: 
	docker build -t sample-project --build-arg PORT=$(PORT) --build-arg SERVER_CLIENT=$(SERVER_CLIENT) --build-arg SERVER_TOKEN=$(SERVER_TOKEN) --build-arg SERVER_CLIENT_UUID=$(SERVER_CLIENT_UUID) .	
	docker run -p$(PORT):$(PORT) sample-project