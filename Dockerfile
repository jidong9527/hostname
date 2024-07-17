FROM golang:alpine AS builder
COPY main.go /go
RUN CGO_ENABLED=0 GO111MODULE="off" go build -o main

FROM centos:7
RUN yum -y install net-tools bind-utils mysql telnet nc openssh
WORKDIR /
COPY --from=builder /go/main /main
EXPOSE 80
ENTRYPOINT ["/main"]
