# go-dockerizing-minimal-ex

## Dockerfile

```Dockerfile
FROM golang:latest
RUN mkdir /go/src/go-dockerizing-minimal-ex
ADD ./src /go/src/go-dockerizing-minimal-ex
WORKDIR /go/src/go-dockerizing-minimal-ex
ENV PORT=3001
RUN go build -o main .
CMD ["/go/src/go-dockerizing-minimal-ex/main"]
```

Docker Console:

```
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
```

```
$ docker build -t loggar/go-dockerizing-minimal-ex .

$ docker images
REPOSITORY                         TAG                 IMAGE ID            CREATED             SIZE
loggar/go-dockerizing-minimal-ex   latest              a048435ac65d        2 minutes ago       822MB
```

```
$ docker run -p 3030:3001 -it loggar/go-dockerizing-minimal-ex
```

- binding host port 3030 to container’s internal port 3001 (-p 3030:3001)
- displaying stdout output in current terminal (-it)

## docker-compose

`docker-compose.yml`

```yml
version: "2"
services:
  go-dockerizing-minimal-ex:
    build: ./go-dockerizing-minimal-ex
    ports:
       - "3030:3001"
```

```
$ docker-compose up
```