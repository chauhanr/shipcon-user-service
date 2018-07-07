build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/chauhanr/shipcon-user-service proto/user/user.proto
	GOOS=linux GOARCH=amd64 go build -o shipcon-user-service
	docker build -t shipcon-user-service .
	go clean


run:
	docker run -d --net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
        shipcon-user-service