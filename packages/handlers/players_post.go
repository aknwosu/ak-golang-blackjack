package handlers

import (
	"blackjack/packages/platform/games"
	"blackjack/packages/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type playerStruct struct {
	Player1 string
	Player2 string
}

func StartGame(deck *games.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {

		requestBody := playerStruct{}
		c.Bind(&requestBody)

		deck.CardDeck = services.GetDeck()
		result := services.CreateGame(deck.CardDeck, requestBody.Player1, requestBody.Player2)
		c.JSON(http.StatusOK, result)
	}
}
