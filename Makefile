GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=bin
NAME=webdog

# mac
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)/$(NAME)-mac
# windows
build-linux:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)/$(NAME)-linux
# linux
build-win:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)/$(NAME)-win.exe
# all
build-all:
	make build-mac
	make build-win
	make build-linux