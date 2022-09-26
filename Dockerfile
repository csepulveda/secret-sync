FROM golang:alpine3.16 as compile
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./. ./
RUN go build -ldflags "-s -w" -o ./main ./main.go

FROM alpine:3.16
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=compile /app/main /app/main
ENTRYPOINT ["./main"]