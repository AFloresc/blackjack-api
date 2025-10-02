package game

import "blackjack-api/models"

type GameSession struct {
	DeckManager *DeckManager
	Player      models.Player
	Dealer      models.Dealer
	GameOver    bool
	Winner      string
}

func NewGameSession() *GameSession {
	deck := NewDeckManager()
	player := models.Player{Hand: []models.Card{deck.DrawCard(), deck.DrawCard()}}
	dealer := models.Dealer{Hand: []models.Card{deck.DrawCard(), deck.DrawCard()}}

	return &GameSession{
		DeckManager: deck,
		Player:      player,
		Dealer:      dealer,
		GameOver:    false,
		Winner:      "",
	}
}

func (gs *GameSession) Hit() {
	gs.Player.Hand = append(gs.Player.Hand, gs.DeckManager.DrawCard())
	if IsBust(gs.Player.Hand) {
		gs.GameOver = true
		gs.Winner = "dealer"
	}
}

func (gs *GameSession) Stand() {
	for CalculateScore(gs.Dealer.Hand) < 17 {
		gs.Dealer.Hand = append(gs.Dealer.Hand, gs.DeckManager.DrawCard())
	}
	gs.GameOver = true
	playerScore := CalculateScore(gs.Player.Hand)
	dealerScore := CalculateScore(gs.Dealer.Hand)

	if IsBust(gs.Dealer.Hand) || playerScore > dealerScore {
		gs.Winner = "player"
	} else if dealerScore > playerScore {
		gs.Winner = "dealer"
	} else {
		gs.Winner = "tie"
	}
}

func (gs *GameSession) GetState() models.GameState {
	return models.GameState{
		Deck:        gs.DeckManager.deck,
		Player:      gs.Player,
		Dealer:      gs.Dealer,
		PlayerScore: CalculateScore(gs.Player.Hand),
		DealerScore: CalculateScore(gs.Dealer.Hand),
		GameOver:    gs.GameOver,
		Winner:      gs.Winner,
		PlayerBust:  IsBust(gs.Player.Hand),
		DealerBust:  IsBust(gs.Dealer.Hand),
	}
}
