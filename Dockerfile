FROM golang:1.9
MAINTAINER Eirik Martiniussen Sylliaas <eirik@sylliaas.no>

RUN go get -u github.com/golang/dep/cmd/dep
RUN mkdir -p /go/src/github.com/eirsyl/statuspage

COPY . /go/src/github.com/eirsyl/statuspage/
WORKDIR /go/src/github.com/eirsyl/statuspage

RUN set -e \
    && dep ensure \
    && go build

CMD ["./statuspage"]