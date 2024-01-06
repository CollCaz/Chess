package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type Game struct {
	Event  string
	Site   string
	Date   time.Time
	Round  int8
	White  string
	Black  string
	Result string
	// PGN    [][2]string
	PGN string
}

func (g Game) MarshalJSON() ([]byte, error) {
	aux := struct {
		Event  string `json:"event"`
		Site   string `json:"site"`
		Date   string `json:"date"`
		Round  int8   `json:"round"`
		White  string `json:"white"`
		Black  string `json:"black"`
		Result string `json:"result"`
		PGN    string `json:"pgn"`
	}{
		Event:  g.Event,
		Site:   g.Site,
		Date:   g.Date.Format(time.DateOnly),
		Round:  g.Round,
		White:  g.White,
		Black:  g.Black,
		Result: g.Result,
		PGN:    fmt.Sprint(g.PGN),
	}

	return json.Marshal(aux)
}

func GetGame(id int) (*Game, error) {
	game := Game{
		Event:  "Chess Open",
		Site:   "Moscow",
		Date:   time.Now(),
		Round:  5,
		White:  "Carlson, Magnus",
		Black:  "Neiman, hans",
		Result: "1-0",
		PGN:    "1. d4 e6 2. Nf3 f5",
	}
	return &game, nil
}
