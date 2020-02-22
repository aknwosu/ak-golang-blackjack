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

func StartGame(blkjack *games.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := playerStruct{}
		c.Bind(&body)

		blkjack.CardDeck = services.GetDeck()

		result := services.CreateGame(blkjack.CardDeck, body.Player1, body.Player2)
		c.JSON(http.StatusOK, result)
	}
}
