# build stage
FROM golang:1.17-alpine3.14 AS builder
RUN mkdir 2chance
RUN chmod 775 -R ./2chance
COPY . /2chance
WORKDIR /2chance
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o 2chance ./cmd/api/main.go

# final stage
FROM alpine:latest AS final
RUN mkdir app
RUN chmod 775 -R ./app
COPY --from=builder /2chance /app
WORKDIR /app
ENV TZ=Asia/Almaty
RUN chmod +x ./2chance
RUN apk add --no-cache tzdata
ENTRYPOINT ["./dictionary-rn"]
EXPOSE 3000
