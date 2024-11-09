FROM golang:1.21 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server 
# ENTRYPOINT [ "/app/server" ]

FROM scratch
RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates
COPY --from=builder /app/server /server
COPY --from=builder /app/public /public 
ENTRYPOINT [ "/server" ]

