# generate stubs/skels
build:
	go build -o bggserver server.go util.go
	go build -o bggclient client.go util.go
	go build -o bgggateway gateway.go util.go

gen-gateway:
	# generate the stub/skel
	protoc messages/bgg.proto \
		-I. \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--go_out=plugins=grpc:.
	# generate the gateway
	protoc messages/bgg.proto \
		-I. \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:.

gen:
	protoc messages/bgg.proto --go_out=plugins=grpc:.

run-server:
	go run server.go util.go

run-client:
	go run client.go util.go $(GAMEID)

run-gateway:
	go run gateway.go util.go 
