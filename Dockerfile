FROM  golang:1.13.9-stretch

WORKDIR $GOPATH/src/legato_server

COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
