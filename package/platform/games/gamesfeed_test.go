package games

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	feed := NewDeck()

	if len(feed.CardDeck) != 0 {
		t.Errorf("CardDeck Repo not initialized successfully")
	}
}
