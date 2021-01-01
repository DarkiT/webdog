GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=bin
NAME=webdog

# mac
build-mac:
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)/$(NAME)-mac
# linux
build-linux:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)/$(NAME)-linux
# windows
build-win:
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)/$(NAME)-win.exe
# all
build-all:
	make build-mac
	make build-win
	make build-linux