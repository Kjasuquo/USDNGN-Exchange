FROM golang:1.18-alpine3.14 as build

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/main.go

# Now copy it into our base image.
FROM alpine:3.13
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 8085
CMD ["/app/main"]

