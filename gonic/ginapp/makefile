start:
	go run .
	#arun sh -c 'swag init -g pkg/main.go && go run ./pkg'
	#sh -c 'swag init --parseDependency --parseInternal --parseDepth 1 && go run .'

init: 
	#go install github.com/swaggo/swag/cmd/swag@latest -insecure
	#go install github.com/swaggo/swag/cmd/swag@v1.8.1
	go install github.com/swaggo/swag/cmd/swag@latest
install:
	# go install .
	go install -ldflags="-s -w -X ginapp/conf.BuildDate=$(shell date -Iseconds) -X ginapp/conf.BuildBranch=$(shell git rev-parse --abbrev-ref HEAD)" 
	#go build -o ~/bin/ginapp .
build:
	go build -ldflags="-s -w -X ginapp/conf.BuildDate=$(shell date -Iseconds) -X ginapp/conf.BuildBranch=$(shell git rev-parse --abbrev-ref HEAD)" -o ginapp main.go
generate:
	#//go:generate swag init --parseDependency --parseInternal -g description.go -o ./docs
	go generate
doc: 
	# swag init -g cmd/main.go
	# swag init --parseDependency --parseInternal --parseDepth 1 #&& open http://m:4500/swagger/index.html
	swag init  #&& open http://m:4500/swagger/index.html

benchcpu:
	go-wrk  -d=50 -c=50  http://localhost:4500/cpu/5

profcpu:
	go tool pprof http://127.0.0.1:4500/debug/pprof/profile

docker:
	docker build -t ginapp .
	docker run -p 4502:4501 --rm -it  ginapp ./main -p 4501

ALL_SRC := $(shell find . -name "*_test.go" | grep -v \
	-e ".*/\..*" \
	-e ".*/_.*" \
	-e ".*/mocks.*")
TEST_DIRS := $(sort $(dir $(filter %_test.go,$(ALL_SRC))))

COVERAGE_FILE := cover.out
.PHONY: test
test:
	echo $(TEST_DIRS)
	@rm -f test.log
	@rm -f $(COVERAGE_FILE)
	@for dir in $(TEST_DIRS); do \
		go test -timeout 20m -coverprofile="cover.tmp" "$$dir" | tee -a test.log; \
		cat cover.tmp >> $(COVERAGE_FILE); \
	done;
	@rm -f cover.tmp
