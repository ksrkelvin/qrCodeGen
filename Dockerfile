# Etapa de construção
FROM golang:1.21 as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# Instala os certificados de CA no builder
RUN apt-get update && apt-get install -y ca-certificates

# Etapa final usando 'scratch'
FROM scratch

# Copia o binário compilado e os certificados raiz para o contêiner final
COPY --from=builder /app/server /server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/public /public

ENTRYPOINT [ "/server" ]