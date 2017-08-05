FROM golang

Add . /go/src/github.com/alwindoss/robin
RUN go get github.com/alwindoss/robin/...
RUN go install github.com/alwindoss/robin
ENTRYPOINT /go/bin/robin start server
EXPOSE 8080