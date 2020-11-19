FROM golang:buster

COPY dockerscript/build.sh /usr/bin

ENTRYPOINT [ "/usr/bin/build.sh" ]