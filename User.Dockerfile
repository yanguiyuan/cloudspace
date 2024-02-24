FROM golang:1.21.7-alpine3.18
LABEL authors="yanguiyuan"
WORKDIR /go/src
COPY cmd cmd
COPY config config
COPY internal internal
COPY pkg pkg
COPY go.mod go.mod
RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod tidy
CMD ["go","run","/go/src/cmd/user"]
