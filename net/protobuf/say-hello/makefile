dev:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	brew install protoc
gen:
	protoc --go_opt=paths=source_relative --go_out=. helloworld.proto
#this will generate helloworld.pb.go
