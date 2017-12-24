FROM ethereum/solc:stable


# Install Go
RUN apk update && apk add curl git bash mercurial build-base linux-headers bzr 'go=1.8.4-r0' && rm -rf /var/cache/apk/*

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin


# Install go-ethereum
RUN go get github.com/ethereum/go-ethereum          &&\
    cd $GOPATH/src/github.com/ethereum/go-ethereum/ &&\
    make                                            &&\
    make devtools

ENTRYPOINT []