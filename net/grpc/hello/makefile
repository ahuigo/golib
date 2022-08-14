#this will generate helloworld.pb.go helloworld_grpc.pb.go
gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pb/helloworld.proto
dev:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	brew install protoc
test:
	curl -v 'http://localhost:8052/rest/n/thinkerou'
