FROM debian:testing

RUN apt-get update && apt-get install -y \
    git \
    libgit2-dev \
    golang \
    texlive \
    texlive-lang-all \
    pandoc \
 && rm -rf /var/lib/apt/lists/*

 ENV GOPATH /go
 ENV PATH $PATH:$GOPATH/bin

 RUN go get github.com/Dhiver/layoutForMe

 RUN go install github.com/Dhiver/layoutForMe

 ENTRYPOINT /bin/bash
