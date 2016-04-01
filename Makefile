export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


default:
	@echo "\033[92mRunning Server ...\033[0m"
	@go run server.go url2rst.go buy123.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go

install:
	@echo "\033[92mInstalling packages ...\033[0m"
	go get -u github.com/PuerkitoBio/goquery
