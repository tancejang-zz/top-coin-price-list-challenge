#We specify our workdir
ARG mypath=/go/src/github.com/tancejang/top-coin-price-list-challenge

# We specify the base image we need for our
# go application
FROM golang
# We copy everything in the root directory
# into our /app directory

ADD . /go/src/github.com/tancejang/top-coin-price-list-challenge/pricing-service
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /go/src/github.com/tancejang/top-coin-price-list-challenge/pricing-service

# Dependencies list to be installed
RUN go get -v gopkg.in/resty.v0

# we run go build to compile the binary
# executable of our Go program
RUN go build -o main
# Our start command which kicks off
# our newly created binary executable
CMD ["/go/src/github.com/tancejang/top-coin-price-list-challenge/pricing-service/main"]