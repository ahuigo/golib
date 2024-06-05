server:
	go test -v server_test.go
client:
	go test -v client_test.go
#this will generate helloworld.pb.go helloworld_grpc.pb.go
gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pb/helloworld.proto
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	brew install protoc
