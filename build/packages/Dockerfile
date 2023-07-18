ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /usr/app

COPY . .

RUN go mod download

ARG ADAPTER=http/fiber

ENV CMD_ADAPTER=${ADAPTER}
RUN go build -o /go/bin/app cmd/$CMD_ADAPTER/main.go

FROM scratch

WORKDIR /usr/app

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]