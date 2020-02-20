package main

import (
	"blackjack/httpd/handlers"
	"blackjack/platform/games"

	"github.com/gin-gonic/gin"
)

func main() {
	deck := games.NewDeck()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server running",
		})
	})
	r.POST("/start-game", handlers.StartGame(deck))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
