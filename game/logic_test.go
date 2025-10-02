package game

import (
	"blackjack-api/models"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	tests := []struct {
		name     string
		hand     []models.Card
		expected int
	}{
		{"Simple sum", []models.Card{{Name: "5", Value: 5}, {Name: "7", Value: 7}}, 12},
		{"Face cards", []models.Card{{Name: "K", Value: 10}, {Name: "Q", Value: 10}}, 20},
		{"Ace as 11", []models.Card{{Name: "A", Value: 11}, {Name: "8", Value: 8}}, 19},
		{"Ace as 1", []models.Card{{Name: "A", Value: 11}, {Name: "K", Value: 10}, {Name: "5", Value: 5}}, 16},
		{"Multiple Aces", []models.Card{{Name: "A", Value: 11}, {Name: "A", Value: 11}, {Name: "9", Value: 9}}, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			score := CalculateScore(tt.hand)
			if score != tt.expected {
				t.Errorf("got %d, want %d", score, tt.expected)
			}
		})
	}
}

func TestIsBust(t *testing.T) {
	tests := []struct {
		name     string
		hand     []models.Card
		expected bool
	}{
		{"Not bust", []models.Card{{Name: "9", Value: 9}, {Name: "10", Value: 10}}, false},
		{"Bust", []models.Card{{Name: "K", Value: 10}, {Name: "Q", Value: 10}, {Name: "5", Value: 5}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bust := IsBust(tt.hand)
			if bust != tt.expected {
				t.Errorf("got %v, want %v", bust, tt.expected)
			}
		})
	}
}
