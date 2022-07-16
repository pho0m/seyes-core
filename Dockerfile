FROM golang:1.15.1-alpine as base

WORKDIR /app

RUN apk add --update --no-cache git ca-certificates make zip tzdata build-base
RUN \
    cd /tmp && \
    GO111MODULE=on go get github.com/cortesi/modd/cmd/modd && \
    cd -

RUN zip -q -r -0 /zoneinfo.zip /usr/share/zoneinfo


# Build pahse
FROM base as builder

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app

RUN CGO_ENABLED=0 go build -o bin/mns-core-server -a -tags netgo -ldflags '-extldflags "-static"' cmd/mns-core-server/main.go

RUN cp bin/mns-core-server /go/bin/mns-core-server

FROM scratch
WORKDIR /app
ENV ZONEINFO /zoneinfo.zip

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/bin/mns-core-server /mns-core-server

EXPOSE 3000
CMD ["/mns-core-server"]
