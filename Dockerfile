FROM golang:alpine AS builder
COPY main.go /go
RUN CGO_ENABLED=0 GO111MODULE="off" go build -o main

FROM centos:8
RUN mkdir /etc/yum.repos.d/bak; mv /etc/yum.repos.d/*.repo /etc/yum.repos.d/bak ;\
    curl -o /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-8.repo
RUN yum -y install net-tools bind-utils mysql telnet nc openssh openssh-clients

WORKDIR /
COPY --from=builder /go/main /main
EXPOSE 80
ENTRYPOINT ["/main"]
