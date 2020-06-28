FROM golang:1.18.5-alpine AS build

MAINTAINER https://github.com/earth-hotfix/webdavGo

WORKDIR /project/
COPY . /project/

RUN export GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go mod tidy
RUN CGO_ENABLE=0 GO111MODULE=on GOOS=linux go build

FROM alpine:latest

WORKDIR /

COPY --from=build /project/webdavGo /webdavGo

ENTRYPOINT ["./webdavGo"]
