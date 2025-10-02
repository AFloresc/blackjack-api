package models

type Card struct {
	Name  string `json:"name"`  // Ej: "A", "K", "10"
	Suit  string `json:"suit"`  // Ej: "♠", "♥"
	Value int    `json:"value"` // Ej: 11, 10, 2...
}

type Player struct {
	Hand []Card `json:"hand"`
}

type Dealer struct {
	Hand []Card `json:"hand"`
}

type GameState struct {
	Deck        []Card `json:"deck"`
	Player      Player `json:"player"`
	Dealer      Dealer `json:"dealer"`
	PlayerScore int    `json:"playerScore"`
	DealerScore int    `json:"dealerScore"`
	GameOver    bool   `json:"gameOver"`
	Winner      string `json:"winner"` // "player", "dealer", "tie", ""
	PlayerBust  bool   `json:"playerBust"`
	DealerBust  bool   `json:"dealerBust"`
}
