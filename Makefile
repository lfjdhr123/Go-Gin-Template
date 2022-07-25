help: Makefile
	@echo "Hello, there! Please Choose a target command run:"
	@echo
	@echo "  build    		:Compile for local OS and ARCH"
	@echo "  linux    		:Compile for Linux platform on AMD64"
	@echo "  clean   		:Clean up all built binaries"
	@echo


.PHONY: build linux clean dev-pack

build: clean
	@go build -o build/bin/apiserver main.go
	@cp -r ./conf build/bin
linux: clean
	@GOOS=linux GOARCH=amd64 go build -o build/linux-amd64/apiserver main.go
	@mkdir build/linux-amd64/conf/
clean:
	@rm -rf build