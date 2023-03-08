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


# FROM golang:1.15.1-alpine as base

# WORKDIR /app

# RUN apk add --update --no-cache git ca-certificates make zip tzdata build-base
# RUN \
#     cd /tmp && \
#     GO111MODULE=on go get github.com/cortesi/modd/cmd/modd && \
#     cd -

# RUN zip -q -r -0 /zoneinfo.zip /usr/share/zoneinfo

# # Build pahse
# FROM base as builder

# COPY go.mod /app
# COPY go.sum /app

# RUN go mod download

# COPY . /app

# RUN CGO_ENABLED=0 go build -o bin/seyes-core-server -a -tags netgo -ldflags '-extldflags "-static"' cmd/seyes-core-server/main.go

# RUN cp bin/seyes-core-server /go/bin/seyes-core-server

# FROM scratch
# WORKDIR /app
# ENV ZONEINFO /zoneinfo.zip

# COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# COPY --from=builder /go/bin/seyes-core-server /seyes-core-server

# RUN apk add git

# WORKDIR /go/src/app
# COPY . .

# ARG TARGETOS TARGETARCH TARGETVARIANT

# ENV CGO_ENABLED=0
# RUN go get \
#     && go mod download \
#     && GOOS=${TARGETOS} GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT#"v"} go build -a -o rtsp-to-web

# FROM alpine:3.17

# WORKDIR /app

# COPY --from=builder /go/src/app/rtsp-to-web /app/
# COPY --from=builder /go/src/app/web /app/web

# RUN mkdir -p /config
# COPY --from=builder /go/src/app/config.json /config

# ENV GO111MODULE="on"
# ENV GIN_MODE="release"

# EXPOSE 3000

# CMD ["./rtsp-to-web", "--config=/config/config.json", "/seyes-core-server"]
