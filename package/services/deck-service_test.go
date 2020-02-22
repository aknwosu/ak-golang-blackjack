package services

import (
	"reflect"
	"testing"
)

func TestGetDeck(t *testing.T) {
	d := GetDeck()

	if len(d) == 0 {
		t.Errorf("CardDeck Repo not initialized successfully")
	}

	if len(d) != 52 {
		t.Errorf("CardDeck should have 52 cards but got %d", len(d))
	}

	if reflect.TypeOf(d[0].Suit) == nil {
		t.Errorf("Each card should have a suit")
	}
}
