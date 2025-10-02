package game

import (
	"blackjack-api/models"
	"testing"
)

func TestNewGameSession(t *testing.T) {
	session := NewGameSession()

	if session == nil {
		t.Fatal("NewGameSession devolvió nil")
	}
	if len(session.Player.Hand) != 2 {
		t.Errorf("esperado 2 cartas para el jugador, recibido %d", len(session.Player.Hand))
	}
	if len(session.Dealer.Hand) != 2 {
		t.Errorf("esperado 2 cartas para el dealer, recibido %d", len(session.Dealer.Hand))
	}
}

func TestHit(t *testing.T) {
	session := NewGameSession()
	initialLen := len(session.Player.Hand)

	session.Hit()

	if len(session.Player.Hand) != initialLen+1 {
		t.Errorf("esperado %d cartas tras hit, recibido %d", initialLen+1, len(session.Player.Hand))
	}

	if session.GameOver && !IsBust(session.Player.Hand) {
		t.Error("GameOver activado sin bust del jugador")
	}
}

func TestStand(t *testing.T) {
	session := NewGameSession()
	session.Stand()

	if !session.GameOver {
		t.Error("esperado GameOver=true tras stand")
	}
	if session.Winner == "" {
		t.Error("esperado campo Winner definido tras stand")
	}
}

func TestGetState(t *testing.T) {
	session := NewGameSession()
	state := session.GetState()

	if state.PlayerScore != CalculateScore(session.Player.Hand) {
		t.Errorf("playerScore incorrecto: %d vs %d", state.PlayerScore, CalculateScore(session.Player.Hand))
	}
	if state.DealerScore != CalculateScore(session.Dealer.Hand) {
		t.Errorf("dealerScore incorrecto: %d vs %d", state.DealerScore, CalculateScore(session.Dealer.Hand))
	}
	if state.GameOver != session.GameOver {
		t.Error("GameOver no coincide con la sesión")
	}
}

func TestMultipleHits(t *testing.T) {
	session := NewGameSession()
	initialLen := len(session.Player.Hand)

	for i := 0; i < 3; i++ {
		session.Hit()
	}

	if len(session.Player.Hand) != initialLen+3 {
		t.Errorf("esperado %d cartas tras 3 hits, recibido %d", initialLen+3, len(session.Player.Hand))
	}

	if session.GameOver && !IsBust(session.Player.Hand) {
		t.Error("GameOver activado sin bust del jugador")
	}
}

func TestStandDealerStays(t *testing.T) {
	session := NewGameSession()

	// Forzamos una mano del dealer con score >= 17
	session.Dealer.Hand = []models.Card{
		{Name: "10", Value: 10},
		{Name: "7", Value: 7},
	}

	session.Stand()

	if !session.GameOver {
		t.Error("esperado GameOver=true tras stand")
	}
	if session.Winner == "" {
		t.Error("esperado campo Winner definido")
	}
}

func TestTieScenario(t *testing.T) {
	session := NewGameSession()

	session.Player.Hand = []models.Card{
		{Name: "10", Value: 10},
		{Name: "9", Value: 9},
	}

	session.Dealer.Hand = []models.Card{
		{Name: "K", Value: 10},
		{Name: "9", Value: 9},
	}

	session.Stand()

	if session.Winner != "tie" {
		t.Errorf("esperado empate, recibido %s", session.Winner)
	}
}

func TestDealerBust(t *testing.T) {
	session := NewGameSession()

	session.Dealer.Hand = []models.Card{
		{Name: "K", Value: 10},
		{Name: "Q", Value: 10},
		{Name: "5", Value: 5},
	}

	session.Stand()

	if !IsBust(session.Dealer.Hand) {
		t.Error("esperado bust del dealer")
	}
	if session.Winner != "player" {
		t.Errorf("esperado ganador 'player', recibido %s", session.Winner)
	}
}

func TestPlayerBust(t *testing.T) {
	session := NewGameSession()
	session.Player.Hand = []models.Card{
		{Name: "K", Value: 10},
		{Name: "Q", Value: 10},
		{Name: "5", Value: 5},
	}
	session.Hit()

	if !session.GameOver {
		t.Error("esperado GameOver=true tras bust del jugador")
	}
	if session.Winner != "dealer" {
		t.Errorf("esperado ganador 'dealer', recibido %s", session.Winner)
	}
}

func TestDealerWins(t *testing.T) {
	session := NewGameSession()
	session.Player.Hand = []models.Card{
		{Name: "8", Value: 8},
		{Name: "9", Value: 9},
	}
	session.Dealer.Hand = []models.Card{
		{Name: "10", Value: 10},
		{Name: "9", Value: 9},
	}
	session.Stand()

	if session.Winner != "dealer" {
		t.Errorf("esperado ganador 'dealer', recibido %s", session.Winner)
	}
}

func TestExactTie(t *testing.T) {
	session := NewGameSession()
	session.Player.Hand = []models.Card{
		{Name: "10", Value: 10},
		{Name: "7", Value: 7},
	}
	session.Dealer.Hand = []models.Card{
		{Name: "K", Value: 10},
		{Name: "7", Value: 7},
	}
	session.Stand()

	if session.Winner != "tie" {
		t.Errorf("esperado empate, recibido %s", session.Winner)
	}
}

func TestDealerDrawsMultipleCards(t *testing.T) {
	session := NewGameSession()

	// Forzamos una mano débil del dealer
	session.Dealer.Hand = []models.Card{
		{Name: "2", Value: 2},
		{Name: "3", Value: 3},
	}

	session.Stand()

	if CalculateScore(session.Dealer.Hand) < 17 {
		t.Errorf("dealer debería haber robado hasta alcanzar al menos 17, pero tiene %d", CalculateScore(session.Dealer.Hand))
	}
	if !session.GameOver {
		t.Error("esperado GameOver=true tras stand")
	}
}

func TestGetStateWithBust(t *testing.T) {
	session := NewGameSession()
	session.Player.Hand = []models.Card{
		{Name: "K", Value: 10},
		{Name: "Q", Value: 10},
		{Name: "5", Value: 5},
	}
	session.Hit()

	state := session.GetState()

	if !state.PlayerBust {
		t.Error("esperado PlayerBust=true")
	}
	if state.Winner != "dealer" {
		t.Errorf("esperado ganador 'dealer', recibido %s", state.Winner)
	}
}

func TestCalculateScoreWithManyAces(t *testing.T) {
	hand := []models.Card{
		{Name: "A", Value: 11},
		{Name: "A", Value: 11},
		{Name: "A", Value: 11},
		{Name: "9", Value: 9},
	}

	score := CalculateScore(hand)

	if score != 12 {
		t.Errorf("esperado score 12 con múltiples ases, recibido %d", score)
	}
}
