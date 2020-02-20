package handlers

import (
	"blackjack/platform/games"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type result struct {
	Winner  string `json:"winner"`
	Players games.Players
}
type playerStruct struct {
	Player1 string
	Player2 string
}

func getDeck() games.CardDeck {
	resp, err := http.Get("https://teston-backend-case.herokuapp.com/shuffle")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	receivedDeck := games.CardDeck{}
	err = json.Unmarshal(data, &receivedDeck)
	return receivedDeck
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

func drawCards(deck games.CardDeck, p games.Player) (games.CardDeck, games.Player) {
	c, v := getCardVal(deck[:2])
	deck = deck[2:]
	p.Cards = append(p.Cards, c...)
	p.Points += v

	return deck, p
}

func continueGame(deck games.CardDeck, players games.Players, index int) result {
	gr := result{}
	currentPlayer := players[index]
	deck, currentPlayer = drawCards(deck, players[index])
	players[index] = currentPlayer
	if currentPlayer.Points >= 17 {
		if index == 0 {
			gr.Winner = players[1].Name
		}
		gr.Players = players
		return gr
	}
	nextIndex := index + 1
	return continueGame(deck, players, nextIndex)
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

func createGame(deck games.CardDeck, p1 string, p2 string) result {
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

func StartGame(deck *games.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		deck.CardDeck = getDeck()
		requestBody := playerStruct{}
		c.Bind(&requestBody)
		result := createGame(deck.CardDeck, requestBody.Player1, requestBody.Player2)
		c.JSON(http.StatusOK, result)
	}
}
