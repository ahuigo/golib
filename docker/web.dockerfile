FROM golang:1.12.7-alpine3.10 AS build-base
ENV GO111MODULE=on
RUN apk add --no-cache git=2.22.0-r0

FROM build-base AS build-env
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV D=/app

COPY ./ $D/

WORKDIR $D/

RUN go mod download && go mod vendor

RUN go build -mod=vendor -o /bin/data-ops ./cmd/main.go

FROM alpine:3.10.2
RUN apk --no-cache add ca-certificates && \
    apk add --no-cache tzdata
COPY --from=build-env /bin/data-ops /data-ops
COPY ./config/conf.yaml /config/conf.yaml
ENV GIN_MODE=release
CMD ["/data-ops"]
