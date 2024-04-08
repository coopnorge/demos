generate:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	buf generate proto

run-server:
	go run github.com/coopnorge/bendik-connectrpc/cmd/server

grpcurl:
	grpcurl \
	-proto proto/coopnorge/store/v1/store_api.proto  \
	-plaintext \
	-d '{"id": "1"}' \
	localhost:8080 coopnorge.store.v1.StoreInformationAPI/GetStore

curl:
	curl \
    --header "Content-Type: application/json" \
    --data '{"id": "2"}' \
    http://localhost:8080/coopnorge.store.v1.StoreInformationAPI/GetStore

run-client:
	go run github.com/coopnorge/bendik-connectrpc/cmd/client
