FROM golang:1.20.0-buster
ENV GO111MODULE on

WORKDIR /go/app

COPY ../app/src /go/app/

RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    github.com/ramya-rao-a/go-outline \
    github.com/nsf/gocode \
    github.com/acroca/go-symbols \
    github.com/fatih/gomodifytags \
    github.com/josharian/impl \
    github.com/haya14busa/goplay/cmd/goplay \
    github.com/go-delve/delve/cmd/dlv \
    golang.org/x/lint/golint \
    golang.org/x/tools/gopls
