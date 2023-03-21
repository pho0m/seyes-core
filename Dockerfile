FROM golang:1.19-alpine as base

WORKDIR /app

RUN apk add --update --no-cache git ca-certificates make zip tzdata build-base
RUN \
    cd /tmp && \
    GO111MODULE=off go get github.com/cortesi/modd/cmd/modd && \
    cd -

RUN zip -q -r -0 /zoneinfo.zip /usr/share/zoneinfo

# Build pahse
FROM base as builder

COPY go.mod /app
COPY go.sum /app

RUN go mod tidy

COPY . /app

RUN CGO_ENABLED=0 go build -o bin/seyes-core-server -a -tags netgo -ldflags '-extldflags "-static"' cmd/seyes-core-server/main.go

RUN cp bin/seyes-core-server /go/bin/seyes-core-server

FROM scratch
WORKDIR /app
ENV ZONEINFO /zoneinfo.zip

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/bin/seyes-core-server /seyes-core-server

EXPOSE 3000
CMD ["/seyes-core-server"]

