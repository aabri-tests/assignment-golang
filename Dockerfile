FROM golang:1.18.3-buster AS build

WORKDIR /app

COPY . .

# Build the Go application
RUN go mod download
RUN make build

FROM alpine:3.17.2

COPY --from=build /app/server /usr/local/bin/

WORKDIR /

EXPOSE 8080

CMD ["/usr/local/bin/server"]