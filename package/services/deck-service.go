package services

import (
	"blackjack/package/platform/games"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetDeck returns a shuffled set of cards from the required api
func GetDeck() games.CardDeck {
	resp, err := http.Get("https://teston-backend-case.herokuapp.com/shuffle")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	receivedDeck := games.CardDeck{}
	err = json.Unmarshal(data, &receivedDeck)
	return receivedDeck
}
