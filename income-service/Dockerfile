FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /src/cmd/income

RUN go build -o /bin/income

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=build /bin/income /bin/income

CMD ["/bin/user"]
