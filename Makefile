.PHONY: all clean

GOBIN:=go
TARGET:=qpfs

all:
	@$(GOBIN) build -ldflags "-X main.Version=`date -u '+%Y-%m-%d_%I:%M:%S'`"

build_win:
	@GOOS=windows GOARCH=amd64 $(GOBIN) build -ldflags "-X main.Version=`date -u '+%Y-%m-%d_%I:%M:%S'`" 

clean:
	rm -f $(TARGET)