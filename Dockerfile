# 🐹 Imagen base con Go 1.24.3 sobre Alpine para ligereza
FROM golang:1.24.3-alpine

# 🛠️ Instala herramientas necesarias (opcional si usas net/http puro)
RUN apk add --no-cache git

# 📁 Crea directorio de trabajo
WORKDIR /app

# 📦 Copia y descarga dependencias
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 📄 Copia el resto del proyecto
COPY . .

# 🔨 Compila el binario
RUN go build -o blackjack-api .

# 🌐 Expone el puerto de la API
EXPOSE 8080

# 🚀 Comando de arranque
CMD ["./blackjack-api"]
