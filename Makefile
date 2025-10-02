# 🔧 Configuración
BINARY_NAME=blackjack-api
PORT=8080
DOCKER_IMAGE=blackjack-api

# 🧪 Ejecuta la API localmente
run:
	@go run main.go

# 🔨 Compila el binario
build:
	@go build -o $(BINARY_NAME) .

# 🧹 Limpia binarios y caché
clean:
	@go clean
	@rm -f $(BINARY_NAME)

# 🐳 Construye la imagen Docker
docker-build:
	@docker build -t $(DOCKER_IMAGE) .

# 🚀 Ejecuta el contenedor Docker
docker-run:
	@docker run -p $(PORT):8080 $(DOCKER_IMAGE)

# 🧪 Ejecuta en segundo plano
docker-up:
	@docker run -d -p $(PORT):8080 --name blackjack $(DOCKER_IMAGE)

# 🧨 Elimina el contenedor
docker-down:
	@docker rm -f blackjack


#make run           # Ejecuta localmente
#make build         # Compila el binario
#make docker-build  # Construye la imagen Docker
#make docker-run    # Ejecuta el contenedor
#make docker-up     # Ejecuta en segundo plano
#make docker-down   # Elimina el contenedor
