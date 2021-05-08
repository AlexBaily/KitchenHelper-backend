FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache git mercurial build-base \
    && go get -d -v ./... \
    && go install -v ./... \
    && apk del git mercurial


CMD ["KitchenHelper-backend"]