package games

type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

type CardDeck []Card

type Player struct {
	Name   string
	Points int
	Cards  []string
}

type Players []Player

type Repo struct {
	CardDeck []Card
}

func NewDeck() *Repo {
	return &Repo{
		CardDeck: []Card{},
	}
}
