package data

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GameModel struct {
	DB *pgxpool.Pool
}

type Game struct {
	ID        int
	Version   int
	CreatedAt time.Time
	Event     string
	Site      string
	Date      time.Time
	Round     int8
	White     int
	Black     int
	Result    string
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
		Result: g.Result,
		PGN:    fmt.Sprint(g.PGN),
		// White:  g.White,
		// Black:  g.Black,
	}

	return json.Marshal(aux)
}

func (gm GameModel) GetGame(id int) (*Game, error) {
	game := Game{}

	query := `
  SELECT id, version, event, site, date, round, white, black, result, pgn
  FROM games
  WHERE id = $1
  `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := gm.DB.QueryRow(ctx, query, id).Scan(
		&game.ID,
		&game.Version,
		&game.Event,
		&game.Site,
		&game.Date,
		&game.Round,
		&game.White,
		&game.Black,
		&game.Result,
		&game.PGN,
	)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (gm GameModel) InsertGame(game *Game) error {
	query := `
  INSERT INTO games (event, site, date, round, white, black, result, pgn)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
  RETURNING id, created_at, version
  `

	args := []any{
		&game.Event,
		&game.Site,
		&game.Date,
		&game.Round,
		&game.White,
		&game.Black,
		&game.Result,
		&game.PGN,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return gm.DB.QueryRow(ctx, query, args...).Scan(&game.ID, &game.CreatedAt, &game.Version)
}
