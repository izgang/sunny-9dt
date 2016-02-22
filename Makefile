# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


default:
	@go run demo.go
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run server.go

fmt:
	@go fmt *.go
