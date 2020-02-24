package handlers

import (
	"blackjack/package/platform/games"
	"blackjack/package/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type playerStruct struct {
	Player1 string
	Player2 string
}

// StartGame is the handler for the start-game route
func StartGame(blkjack *games.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := playerStruct{}
		c.Bind(&body)
		if body.Player1 == "" {
			body.Player1 = "You"
		}
		if body.Player2 == "" {
			body.Player2 = "Player 2"
		}
		// Get card deck for game
		blkjack.CardDeck = services.GetDeck()

		result := services.CreateGame(blkjack.CardDeck, body.Player1, body.Player2)
		c.JSON(http.StatusOK, result)
	}
}
