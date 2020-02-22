package services

import (
	"blackjack/package/platform/games"
	"reflect"
	"testing"
)

func TestCreateGame(t *testing.T) {

	deck := []games.Card{
		{
			Suit:  "CLUBS",
			Value: "3",
		},
		{
			Suit:  "HEARTS",
			Value: "J",
		},
		{
			Suit:  "HEARTS",
			Value: "3",
		},
		{
			Suit:  "SPADES",
			Value: "5",
		},
		{
			Suit:  "CLUBS",
			Value: "4",
		},
		{
			Suit:  "CLUBS",
			Value: "8",
		},
		{
			Suit:  "HEARTS",
			Value: "2",
		},
		{
			Suit:  "SPADES",
			Value: "6",
		},
		{
			Suit:  "DIAMONDS",
			Value: "2",
		},
	}
	{
		// Create game should play the game, return the winner and the Game details
		d := deck
		res := CreateGame(d, "PL1", "PL2")
		if !reflect.DeepEqual(res.Winner, "PL2") {
			t.Errorf("CreateGame() = %v, want %v", res.Winner, "PL2")
		}

		if !(len(res.Players[0].Cards) >= 2) {
			t.Errorf("Min of 2 cards should be drawn")
		}

		if res.Players[0].Cards[0] != "C3" {
			t.Errorf("The first player should have selected the first card in the deck but got %v", res.Players[0].Cards[0])
		}

		if res.Players[0].Points != 25 {
			t.Errorf("The first player should have selected the first card in the deck but got %v", res.Players[0].Points)
		}
	}

	// Test BlackJack
	{
		d := deck
		d[0].Value = "A"
		res := CreateGame(deck, "PL1", "PL2")
		if !reflect.DeepEqual(res.Winner, "PL1") {
			t.Errorf("CreateGame() = %v, want %v", res.Winner, "PL1")
		}

		d = deck
		d[0].Value = "3"
		d[2].Value = "A"
		d[3].Value = "K"
		res = CreateGame(deck, "PL1", "PL2")
		if !reflect.DeepEqual(res.Winner, "PL2") {
			t.Errorf("CreateGame() = %v, want %v", res.Winner, "PL2")
		}
	}

	//Test Continue Game when there is no winner on first draw
	{
		deck := []games.Card{
			{
				Suit:  "CLUBS",
				Value: "3",
			},
			{
				Suit:  "HEARTS",
				Value: "3",
			},
			{
				Suit:  "HEARTS",
				Value: "3",
			},
			{
				Suit:  "SPADES",
				Value: "5",
			},
			{
				Suit:  "CLUBS",
				Value: "4",
			},
			{
				Suit:  "CLUBS",
				Value: "8",
			},
			{
				Suit:  "HEARTS",
				Value: "7",
			},
			{
				Suit:  "SPADES",
				Value: "5",
			},
			{
				Suit:  "DIAMONDS",
				Value: "2",
			},
			{
				Suit:  "DIAMONDS",
				Value: "2",
			},
		}
		// Player2 wins if their score > 17 && less than Player1s
		// score that must be less than 17
		d := deck
		res := CreateGame(d, "PL1", "PL2")
		if !reflect.DeepEqual(res.Winner, "PL2") {
			t.Errorf("Expected PL2 to win based on card selection %v", res)
		}

		// Player1 wins if Player2s score is > 21
		d = deck
		deck[7].Value = "Q"
		res = CreateGame(d, "PL1", "PL2")
		if !reflect.DeepEqual(res.Winner, "PL1") {
			t.Errorf("Expected PL2 to win based on card selection")
		}

		// Game continues for Player2 till Player2s points more than Player 1
		d = deck

		deck[7].Value = "1"
		res = CreateGame(d, "PL1", "PL2")
		if !reflect.DeepEqual(res.Winner, "PL2") {
			t.Errorf("Expected PL2 to win based on card selection %v", res)
		}

		if !reflect.DeepEqual(len(res.Players[0].Cards), 4) && !reflect.DeepEqual(len(res.Players[1].Cards), 6) {
			t.Errorf("Expected PL1 and PL2 to have drawn 4 and 6 cards respectively, but got %v %v", len(res.Players[0].Cards), len(res.Players[1].Cards))
		}
	}
}
