FROM golang:1.17.3-alpine3.14 AS build

WORKDIR /go/src/docker-plugin-secrets

COPY go.mod go.sum main.go ./

RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod download &&\
    go install -v

#FROM scratch
#
#COPY --from=build "/go/bin/docker-plugin-secrets" "/go/bin/docker-plugin-secrets"

ENTRYPOINT ["/go/bin/docker-plugin-secrets"]
