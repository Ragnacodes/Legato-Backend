FROM  golang:1.15-stretch

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download
