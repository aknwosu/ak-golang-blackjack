package games

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

type CardDeck []Card

type SelectedCards struct {
	Player1Cards []Card
	Player2Cards []string
}

type Player struct {
	Name   string
	Points int
	Cards  []string
}

type Players []Player

type Repo struct {
	CardDeck []Card
}

func GetDeck() CardDeck {
	resp, err := http.Get("https://teston-backend-case.herokuapp.com/shuffle")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	receivedDeck := CardDeck{}
	err = json.Unmarshal(data, &receivedDeck)
	return receivedDeck
}

func NewDeck() *Repo {
	return &Repo{
		CardDeck: []Card{},
	}
}
