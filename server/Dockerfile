FROM golang:1.15


RUN go get github.com/lib/pq

RUN mkdir -p /server

COPY . /server

WORKDIR /server
RUN go build src/server.go
ENTRYPOINT ./server


EXPOSE 3000

HEALTHCHECK --interval=10s --timeout=5m \
    CMD curl -f http://localhost/ || exit 1