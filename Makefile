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
	@docker run -p $(PORT):8080 --name $(DOCKER_IMAGE) $(DOCKER_IMAGE) 

# 🧪 Ejecuta en segundo plano
docker-up:
	@docker run -d -p $(PORT):8080 --name $(DOCKER_IMAGE) $(DOCKER_IMAGE)

# 🧨 Elimina el contenedor
docker-down:
	@docker rm -f blackjack-api

# 🔄 Reconstruye y reinicia el contenedor limpio
refresh:
	@docker rm -f blackjack-api || true
	@docker build -t $(DOCKER_IMAGE) .
	@docker run -d -p $(PORT):8080 --name $(DOCKER_IMAGE) $(DOCKER_IMAGE)

# 🛑 Detiene el contenedor sin eliminarlo
stop:
	@docker stop $(DOCKER_IMAGE)

# ▶️ Inicia el contenedor detenido
start:
	@docker start $(DOCKER_IMAGE)