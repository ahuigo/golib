# Face Recognition Project Makefile
# Go version: 1.24.4
# Dependencies: dlib, jpeg

.PHONY: all test build clean install-deps install-mac install-debian models help

# CGO编译环境变量 - macOS配置
export CGO_CXXFLAGS = -std=c++17 -stdlib=libc++ \
	-I/Library/Developer/CommandLineTools/SDKs/MacOSX15.2.sdk/usr/include/c++/v1 \
	-I/opt/homebrew/Cellar/dlib/20.0/include \
	-I/opt/homebrew/opt/jpeg/include

export CGO_LDFLAGS = -L/opt/homebrew/Cellar/dlib/20.0/lib -ldlib \
	-L/opt/homebrew/opt/jpeg/lib -ljpeg \
	-stdlib=libc++

# 运行特定测试
test:
	go test -v -run TestIsSameFace

test2:
	go test -v -run TestRealHumanFace

# 安装依赖 - macOS
install-mac:
	@echo "Installing dependencies on macOS..."
	brew install dlib jpeg cmake pkg-config

# 安装依赖 - Debian/Ubuntu
install-debian:
	@echo "Installing dependencies on Debian/Ubuntu..."
	sudo apt-get update
	sudo apt-get install -y libdlib-dev libjpeg-dev cmake pkg-config build-essential
