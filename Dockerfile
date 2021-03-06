FROM golang:1.13

WORKDIR /go/src/github.com/aknwosu/ak-golang-blackjack

# Install dependencies

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . /go/src/github.com/aknwosu/ak-golang-blackjack

RUN chmod +x /go/src/github.com/aknwosu/ak-golang-blackjack/main.go

COPY main.go main.go

RUN go build

EXPOSE 8080

ENTRYPOINT ["./blackjack"]
