package games

// Card is defines a expected Card object
type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

// CardDeck is the wrapper array cards object
type CardDeck []Card

// Player is the expected property for game players
type Player struct {
	Name   string   `json:"name"`
	Points int      `json:"points"`
	Cards  []string `json:"cards"`
}

// Players is the wrapper for of Player objects
type Players []Player

//Repo defines model mock
type Repo struct {
	CardDeck []Card
	Players  []Player
}

func NewDeck() *Repo {
	return &Repo{
		CardDeck: []Card{},
	}
}

type Meal struct {
	DaySlot string `json:"daySlot,omitempty"`
	Date    string `json:"date"`
}
