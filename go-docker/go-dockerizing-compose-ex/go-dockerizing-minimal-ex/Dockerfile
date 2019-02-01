FROM golang:latest
RUN mkdir /go/src/go-dockerizing-minimal-ex
ADD ./src /go/src/go-dockerizing-minimal-ex
WORKDIR /go/src/go-dockerizing-minimal-ex
RUN go build -o main .
CMD ["/go/src/go-dockerizing-minimal-ex/main"]