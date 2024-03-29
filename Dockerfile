FROM golang:1.22 as base

WORKDIR /app
COPY go.* .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o vanity .

FROM gcr.io/distroless/static-debian11
COPY --from=base /app/vanity ./app

CMD ["/app"]
