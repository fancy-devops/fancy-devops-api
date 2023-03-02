## docker build -t fancy-devops-api:v1.0.0 .
## build
FROM golang:1.18-alpine3.17 AS build
RUN apk add build-base
COPY . /go/src/app
WORKDIR /go/src/app
RUN go env -w GO111MODULE=on
RUN go mod tidy
RUN go build -o gtc

## run
FROM alpine:3.17
RUN mkdir -p /app
COPY conf/* /app/conf/
WORKDIR /app
COPY --from=build /go/src/app/gtc /app/

ENV PATH $PATH:/app
EXPOSE 7000
CMD ["/app/gtc"]