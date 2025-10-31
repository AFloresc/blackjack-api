
# Blackjack API

Simple and basic blackjack API-Rest backend

```
// Object structure excample:
{
    "deck": [],
    "player": {
        "hand": []
    }
    "dealer": {
        "hand": []
    },
    "playerScore": 8,
    "dealerScore": 0,
    "gameOver": false,
    "winner": "",
    "playerBust": false,
    "dealerBust": false
}
```

## API Reference

#### Service status

```http
  GET /blackjack/api/v1/
```

If servide is running you will get this message\

```
{
  "message": "Blackjack API is running...🚀",
  "status": "ok"
}
```

#### Start Game

```http
  POST /blackjack/api/v1/start
```

Starts the game creating a new deck shuffled\
Deals two cards to the player and to the dealer\
Returns the main state of the game\

#### Hit

```http
  POST /blackjack/api/v1/hit
```
Deals a card to the player\
Recalculates score\
If the player's score is > 21 the sate is updated as a bust and the game ends\
Returns the upodated state of the game\


#### Stand

```http
  POST /blackjack/api/v1/stand
```
The player stands\
The dealer plays it's hand, automasticasly hits cards until score >17 or burst\
The winner field is evaluated and updated\
The game gets checked as finished (gameOver = true)\
Returns the final state od the game game\


#### Restart

```http
  POST /blackjack/api/v1/restart
```
Restarts the current game\
Creates a new deck shuffled\
The hands get cleaned\
Returns the upodated state of the game\


#### State

```http
  GET /blackjack/api/v1//state
```
Get the current state of the Game.\




## Run Locally

Use it's dockerized, you can use make file.

Executhe the API locally

```bash
    go run main.go
```

Compile to get binary

```bash
    go build -o $(BINARY_NAME) .
```

Clean binary cache

```bash
    go clean
    rm -f $(BINARY_NAME)
```

Build Docker image

```bash
    docker build -t $(DOCKER_IMAGE) .
```

Run Docker container

```bash
    docker run -p $(PORT):8080 $(DOCKER_IMAGE)
```

Run Docker container

```bash
    docker run -p $(PORT):8080 $(DOCKER_IMAGE)
```

Remove Docker container

```bash
    docker rm -f blackjack







