FROM golang:1.15-stretch

RUN apt update && apt upgrade -y && \
    make openssh-client

WORKDIR /usr/src/app
COPY . . 

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
