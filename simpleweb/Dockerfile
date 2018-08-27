FROM golang:1.10-alpine3.8

RUN apk add git \
	&& go get github.com/codegangsta/gin \
	&& apk del git

ADD . /src

CMD ["go", "run", "/src/server.go"]
