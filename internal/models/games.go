package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
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
	Event     string `json:"event" query:"event"`
	Site      string `json:"site" query:"site"`
	Date      string `json:"date" query:"date"`
	Round     int8   `json:"round" query:"round"`
	White     string `json:"white" query:"white"`
	Black     string `json:"black" query:"black"`
	Result    string `json:"result" query:"result"`
	PGN       string `json:"pgn" query:"pgn"`
}

type SearchGame struct {
	Game    Game
	Offset  int
	OrderBy string
}

func (s *SearchGame) Validate() error {
	// TODO: validate game
	v := new(Game)
	metaValue := reflect.ValueOf(v).Elem()

	field := metaValue.FieldByName(s.OrderBy)
	if (field == reflect.Value{}) {
		return errors.New("OrderBy Field is not valid")
	}

	return nil
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
		Date:   g.Date,
		Round:  g.Round,
		Result: g.Result,
		White:  g.White,
		Black:  g.Black,
		PGN:    fmt.Sprint(g.PGN),
	}

	return json.Marshal(aux)
}

func (gm GameModel) GetGame(id int) (*Game, error) {
	game := Game{}

	query := `
  SELECT id, created_at, version, event, site, date, round, white, black, result, pgn
  FROM games
  WHERE id = $1
  `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tempDate time.Time

	err := gm.DB.QueryRow(ctx, query, id).Scan(
		&game.ID,
		&game.CreatedAt,
		&game.Version,
		&game.Event,
		&game.Site,
		&tempDate,
		&game.Round,
		&game.White,
		&game.Black,
		&game.Result,
		&game.PGN,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	game.Date = tempDate.Format(time.DateOnly)

	return &game, nil
}

func (gm GameModel) QueryGame(search *SearchGame) (*Game, error) {
	game := Game{}
	query := `
  SELECT id, created_at, version, event, site, date, round, white, black, result, pgn
  FROM games
  WHERE round = $1 AND (result = $2)
  `

	sg := search.Game
	args := []any{sg.Round, sg.Result}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	var tempDate time.Time
	// r := int(sg.Round)
	defer cancel()
	err := gm.DB.QueryRow(ctx, query, args...).Scan(
		&game.ID,
		&game.CreatedAt,
		&game.Version,
		&game.Event,
		&game.Site,
		&tempDate,
		&game.Round,
		&game.White,
		&game.Black,
		&game.Result,
		&game.PGN,
	)
	if err != nil {
		fmt.Println(err, "AAAAAAAAAAA")
		return nil, err
	}

	game.Date = tempDate.Format(time.DateOnly)

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
