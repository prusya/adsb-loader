FROM golang:1.20 as build-env

ADD . /src
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /src/svc

###

FROM ubuntu:latest

RUN apt-get update
RUN apt-get install -y curl

COPY --from=build-env /src/svc /svc

ENTRYPOINT ["/svc"]