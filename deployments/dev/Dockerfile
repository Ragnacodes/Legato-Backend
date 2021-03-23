FROM  golang:1.15-stretch

WORKDIR $GOPATH/src/legato_server

COPY go.mod .
RUN go get -d -v ./...

COPY . .
RUN go install -v ./...
