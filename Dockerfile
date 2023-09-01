FROM golang:1.20.1 as base

WORKDIR /go/src/app
COPY go.* .
RUN go mod download

COPY . .
RUN go build -o /go/bin/app ./cmd/gt

FROM gcr.io/distroless/static-debian11
COPY --from=base /go/bin/app /app

CMD ["/app"]
