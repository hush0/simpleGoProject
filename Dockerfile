FROM golang:latest
MAINTAINER shuaihuang217837@sohu-inc.com
WORKDIR $GOPATH/src/simpleGoProject
ADD . $GOPATH/src/simpleGoProject
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io/
RUN go build .
EXPOSE 9090
ENTRYPOINT ["./main"]



FROM private-registry.sohucs.com/domeos-pub/golang:1.14-centos-git
COPY ./ /
RUN go env -w GOPROXY="https://goproxy.cn,direct" && cd /code/go/ && go build -o app main/main.go
WORKDIR /code/go
COPY dumb-init /usr/bin/dumb-init
ENTRYPOINT ["dumb-init"]
CMD ./app