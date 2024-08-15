FROM golang:1.22.2-alpine AS build
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main .
RUN chmod +x ./main
COPY config.json .
EXPOSE 3000
CMD ["./main"]