FROM golang:1.23.0-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /src/cmd/reporting

RUN go build -o /bin/reporting

FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=build /bin/reporting /bin/reporting

CMD ["/bin/user"]
