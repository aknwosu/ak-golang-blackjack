# BlackJack (Go)

This game simulates playing a blackjack game where players draw cards from a deck until a winner who meets the specified requirements is gotten.

## Running the application
  1. In your go directory, clone the application 
  
      `git clone https://github.com/aknwosu/ak-golang-blackjack.git`

  2. Set the repo as current folder
  
      `cd https://github.com/aknwosu/ak-golang-blackjack.git`

  3. Install dependencies
  
      `go mod download`

  4.  Run the application

      `make run`

  
  Other commands available

    Run tests using "make test" or "make test-cover"

    Run in dev mode "make dev" (uses gin to forward port to 3000)

    Build using docker "docker build -t ak-golang-blackjack ."

    Run the docker image "docker run -P ak-golang-blackjack"

## Endpoints
When the server is running, the games endpoint is available at
`POST "localhost:8080/start-game"`
or the otherwise specified port. 

The route optionally takes a body to indicate player names otherwise default names 2 and Player 2 are used.

    {   
        "player1": "First-player-name",
        "player2": "Second-player-name"
    }


## Future Considerations
  If this was to be built on a larger scale, I would have it in a live mode where the first user joins the game and has to wait for another user to join the same room. The first user can then decide to start the game. It could easily be used as a coin flipping result between friends.

  I would also probably have it in a pod and kubernetes cluster using the already created Docker file.

## Requirements
1. Each player takes two cards each from the top of a randomly shuffled deck  provided by our endpoint (https://teston-backend-case.herokuapp.com/shuffle)
  
    * You take the first two cards, Bob takes the next two

2. Calculate the total sum of each players cards
    * Numbered cards have the value on the card
    * Jack (J), Queen (Q) and King (K) each counts as 10 points
    * Ace (A) counts as 11 points
3. If either player has 21 points - blackjack - that player wins the round
4. Otherwise continue drawing cards following the following rules:
    * You draw cards first until
        - the sum of your cards is 17 or higher
        - if you exceed 21 points you loose without Bob having to draw more cards
    * Bob draws cards until
        - the sum of his cards is higher than yours
        - if his cards exceed 21 points he looses

