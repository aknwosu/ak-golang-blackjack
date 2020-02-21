package services

import (
	"blackjack/packages/platform/games"
	"fmt"
	"strconv"
)

type result struct {
	Winner  string `json:"winner"`
	Players games.Players
}

func getCardVal(sc games.CardDeck) ([]string, int) {
	var val int
	cardValue := 0
	var selCd []string
	for i := 0; i < 2; i++ {
		switch sc[i].Value {
		case "A":
			val = 11
		case "J":
			val = 10
		case "Q":
			val = 10
		case "K":
			val = 10
		default:
			val, _ = strconv.Atoi((sc[i].Value))
		}
		selCd = append(selCd, string(sc[i].Suit[0])+string(sc[i].Value[0]))
		cardValue += val
	}
	return selCd, cardValue
}

func drawCards(deck games.CardDeck, p games.Player) (games.CardDeck, games.Player) {
	c, v := getCardVal(deck[:2])
	deck = deck[2:]
	p.Cards = append(p.Cards, c...)
	p.Points += v

	return deck, p
}

func playCards(deck games.CardDeck, players games.Players) (games.CardDeck, games.Players) {
	for i := 0; i < len(players); i++ {
		c, v := getCardVal(deck[:2])
		deck = deck[2:]
		fmt.Println(c, v)
		players[i].Cards = append(players[i].Cards, c...)
		players[i].Points += v
	}
	fmt.Printf("%+v \n", players)
	return deck, players
}

func simpleLoop(deck games.CardDeck, player games.Player, min int) (games.CardDeck, games.Player) {
	deck, player = drawCards(deck, player)
	if player.Points > 21 {
		return deck, player
	}
	if player.Points < min {
		return simpleLoop(deck, player, min)
	}
	return deck, player
}

func CreateGame(deck games.CardDeck, p1 string, p2 string) result {
	player1 := games.Player{Name: p1}
	player2 := games.Player{Name: p2}
	players := games.Players{}
	players = append(players, player1)
	players = append(players, player2)
	deck, players[0] = drawCards(deck, players[0])
	deck, players[1] = drawCards(deck, players[1])
	gameResult := result{}
	gameResult.Players = players

	if players[0].Points == 21 {
		gameResult.Winner = players[0].Name
		return gameResult
	}
	if players[1].Points == 21 {
		gameResult.Winner = p2
		return gameResult
	}

	deck, players[0] = simpleLoop(deck, players[0], 17)
	gameResult.Players = players
	if players[0].Points > 21 {
		gameResult.Winner = p2
		return gameResult
	}
	deck, players[1] = simpleLoop(deck, players[1], players[0].Points)
	gameResult.Players = players
	if players[1].Points > 21 {
		gameResult.Winner = p1
		return gameResult
	}
	gameResult.Players = players
	gameResult.Winner = p2
	return gameResult
}
