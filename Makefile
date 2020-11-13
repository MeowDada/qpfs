.PHONY: all clean

GOBIN:=go
TARGET:=qpfs

all:
	@$(GOBIN) build -ldflags "-X main.version=`git rev-parse HEAD`"

build_win:
	@GOOS=windows GOARCH=amd64 $(GOBIN) build -ldflags "-X main.version=`git rev-parse HEAD`" 

clean:
	rm -f $(TARGET)