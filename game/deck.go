package game

import (
	"blackjack-api/models"
	"math/rand"
	"strconv"
	"time"
)

type DeckManager struct {
	rng  *rand.Rand
	deck []models.Card
}

// Constructor
func NewDeckManager() *DeckManager {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &DeckManager{
		rng:  rng,
		deck: generateDeck(rng),
	}
}

// Genera y baraja el mazo
func generateDeck(rng *rand.Rand) []models.Card {
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

	rng.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}

func (dm *DeckManager) DrawCard() models.Card {
	if len(dm.deck) == 0 {
		dm.deck = generateDeck(dm.rng) // Reinicia si se agota
	}
	card := dm.deck[0]
	dm.deck = dm.deck[1:]
	return card
}
