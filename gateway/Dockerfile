FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /src/cmd/gateway

RUN go build -o /bin/gateway

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=build /bin/gateway /bin/gateway

CMD ["/bin/user"]
