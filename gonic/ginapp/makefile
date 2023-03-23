start:
	#arun sh -c 'swag init -g pkg/main.go && go run ./pkg'
	arun sh -c 'swag init --parseDependency --parseInternal --parseDepth 1 && go run .'

install:
	go install .
	#go build -o ~/bin/ginapp .

generate:
	#//go:generate swag init --parseDependency --parseInternal -g description.go -o ./docs
	go generate
init: doc
	#go install github.com/swaggo/swag/cmd/swag@latest -insecure
	go install github.com/swaggo/swag/cmd/swag@v1.8.1
doc: 
	# swag init -g cmd/main.go
	swag init --parseDependency --parseInternal --parseDepth 1 #&& open http://m:4500/swagger/index.html

benchcpu:
	go-wrk  -d=50 -c=50  http://localhost:4500/cpu/5

profcpu:
	go tool pprof http://127.0.0.1:4500/debug/pprof/profile

docker:
	docker build -t ginapp .
	docker run -p 4502:4501 --rm -it  ginapp ./main -p 4501

ALL_SRC := $(shell find . -name "*_test.go" | grep -v -e vendor \
	-e ".*/\..*" \
	-e ".*/_.*" \
	-e ".*/mocks.*")
TEST_DIRS := $(sort $(dir $(filter %_test.go,$(ALL_SRC))))
COVERAGE_FILE := coverage.out
test:
	echo $(TEST_DIRS)
	@rm -f test.log
	@rm -f $(COVERAGE_FILE)
	@for dir in $(TEST_DIRS); do \
		go test -timeout 20m -coverprofile="test.temp" "$$dir" | tee -a test.log; \
		cat test.temp >> $(COVERAGE_FILE); \
	done;
	@rm -f test.temp
