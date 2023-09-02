FROM golang:1.20.1 as base

WORKDIR /app
COPY go.* .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o gt ./cmd/gt

FROM gcr.io/distroless/static-debian11
COPY --from=base /app/gt ./app

CMD ["/app"]
