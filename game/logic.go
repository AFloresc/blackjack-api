package game

import (
	"blackjack-api/models"
	"math/rand"
	"strconv"
	"time"
)

func NewDeck() []models.Card {
	suits := []string{"♠", "♥", "♦", "♣"}
	names := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	var deck []models.Card

	for _, suit := range suits {
		for _, name := range names {
			var value int
			switch name {
			case "A":
				value = 11
			case "K", "Q", "J":
				value = 10
			default:
				val, _ := strconv.Atoi(name)
				value = val
			}
			deck = append(deck, models.Card{
				Name:  name,
				Suit:  suit,
				Value: value,
			})
		}
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

func StartGame() models.GameState {
	deck := NewDeckManager()
	playerHand := []models.Card{deck.DrawCard(), deck.DrawCard()}
	dealerHand := []models.Card{deck.DrawCard(), deck.DrawCard()}

	return models.GameState{
		Deck:        deck.deck,
		Player:      models.Player{Hand: playerHand},
		Dealer:      models.Dealer{Hand: dealerHand},
		PlayerScore: CalculateScore(playerHand),
		DealerScore: CalculateScore(dealerHand),
	}
}

func CalculateScore(hand []models.Card) int {
	score := 0
	aces := 0

	for _, card := range hand {
		score += card.Value
		if card.Name == "A" {
			aces++
		}
	}

	for score > 21 && aces > 0 {
		score -= 10
		aces--
	}

	return score
}

func IsBust(hand []models.Card) bool {
	return CalculateScore(hand) > 21
}
