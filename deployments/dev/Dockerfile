FROM  golang:1.13.9-stretch

WORKDIR $GOPATH/src/legato_server

COPY ../go.mod .
RUN go get -d -v ./...

COPY ../.. .
RUN go install -v ./...
