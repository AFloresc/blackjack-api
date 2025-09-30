
# Blackjack API

Simple and basic blackjack API-Rest backend
## API Reference

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

**Needs a body object with the previous state.**

#### Stand

```http
  POST /blackjack/api/v1/stand
```
The player stands\
The dealer plays it's hand, automasticasly hits cards until score >17 or burst\
The winner field is evaluated and updated\
The game gets checked as finished (gameOver = true)\
Returns the final state od the game game\

**Needs a body object with the previous state.**

#### Restart

```http
  POST /blackjack/api/v1/restart
```
Restarts the current game\
Creates a new deck shuffled\
The hands get celaned\
Returns the upodated state of the game\

**Needs a body object with the previous state.**

#### State

```http
  GET /blackjack/api/v1//state
```
Get the current state of the Game.\

**Needs a body object with the previous state.**






