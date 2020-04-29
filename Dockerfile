FROM golang:1.14.1-alpine

ENV DB_USER = postgres
ENV DB_PASS = secret
ENV DB_HOST = localhost
ENV DB_NAME = collectionsdb

EXPOSE 8000
RUN apk add --no-cache --update git gcc musl-dev; \
    mkdir -p ${GOPATH}/src/github.com/codefornl/collections-api;
    

WORKDIR ${GOPATH}/src/github.com/codefornl/collections-api
COPY . ${GOPATH}/src/github.com/codefornl/collections-api
RUN go mod vendor && go build -o collections-api .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /go/src/github.com/codefornl/collections-api/collections-api .
CMD [ "/app/collections-api" ]
