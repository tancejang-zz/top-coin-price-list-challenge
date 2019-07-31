FROM golang

RUN mkdir -p /go/src/github.com/tancejang/go_challenge_pricing_service

ADD . /go/src/github.com/tancejang/go_challenge_pricing_service

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/tancejang/go_challenge_pricing_service/pricing/cmd  -watch github.com/tancejang/go_challenge_pricing_service/pricing
