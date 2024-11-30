# doc: post/go/go-debug-vscode.md#docker debug
image:
	docker build -t micro-service-example .
stop:
	docker stop server8
test:
	echo docker run --rm --name server8  --security-opt=seccomp:unconfined -v $$PWD:/build -p8080:8080 -p2345:2345 delve-docker-vscode-example
run:
	docker run --rm --name server8  --security-opt=seccomp:unconfined -v $$PWD:/build -p8080:8080 -p2345:2345 delve-docker-vscode-example
