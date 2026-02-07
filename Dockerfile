# Stage 1: Build (Compilación)
FROM golang:1.24-alpine AS builder

# Instalamos dependencias necesarias para compilar (certificados y tzdata)
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Aprovechamos el cache de capas de Docker para las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el código fuente y las migraciones
COPY . .

# Compilamos el binario
# - CGO_ENABLED=0 para que sea estático (clave para Alpine y ARM)
# - ldflags "-s -w" para reducir el peso eliminando debug info
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app ./cmd/api/main.go

# Stage 2: Final (Imagen ligera de producción)
FROM alpine:latest

# Añadimos certificados (HTTPS) y zona horaria (para logs en hora local)
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copiamos el binario y la carpeta de migraciones desde el builder
COPY --from=builder /app/app .
COPY --from=builder /app/migrations ./migrations

# Usamos una variable de entorno con valor por defecto
ENV PORT=8080
EXPOSE ${PORT}

# Ejecutamos el binario directamente
CMD [ "./app" ]