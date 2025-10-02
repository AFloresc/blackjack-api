package game

import (
	"testing"
)

func TestDrawCard(t *testing.T) {
	dm := NewDeckManager()
	initialLen := len(dm.deck)
	card := dm.DrawCard()

	if len(dm.deck) != initialLen-1 {
		t.Errorf("expected deck size %d, got %d", initialLen-1, len(dm.deck))
	}
	if card.Name == "" || card.Suit == "" || card.Value == 0 {
		t.Errorf("invalid card drawn: %+v", card)
	}
}
