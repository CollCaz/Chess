package data

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PlayerModel struct {
	DB pgxpool.Pool
}

type Player struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdat"`
	WorldRank      int       `json:"worldrank"`
	BYear          int       `json:"byear"`
	Federation     string    `json:"federation"`
	Sex            string    `json:"sex"`
	FideID         int       `json:"fideid"`
	FideTttle      string    `json:"fidetttle"`
	StandardRating int       `json:"standardrating"`
	RapidRating    int       `json:"rapidrating"`
	BlitzRating    int       `json:"blitzrating"`
}

func (p Player) MarshalJSON() ([]byte, error) {
	rating := make(map[string]int)
	rating["Standard"] = p.StandardRating
	rating["Rapid"] = p.RapidRating
	rating["Blitz"] = p.BlitzRating

	aux := struct {
		Name       string
		WorldRank  int
		BYear      int
		Federation string
		Sex        string
		FideID     int
		FideTttle  string
		Rating     map[string]int
	}{
		Name:       p.Name,
		WorldRank:  p.WorldRank,
		BYear:      p.BYear,
		Federation: p.Federation,
		Sex:        p.Sex,
		FideID:     p.FideID,
		FideTttle:  p.FideTttle,
		Rating:     rating,
	}

	return json.Marshal(aux)
}

func (pm *PlayerModel) GetPlayer(id int) (*Player, error) {
	player := Player{
		Name:           "Carlson, Magnus",
		WorldRank:      1,
		BYear:          1990,
		Federation:     "Norway",
		Sex:            "Male",
		FideID:         1503014,
		FideTttle:      "Grandmaster",
		StandardRating: 2830,
		RapidRating:    2823,
		BlitzRating:    2886,
	}
	return &player, nil
}

func (pm *PlayerModel) InsertPlayer(p *Player) error {
	query := `
  INSERT INTO players (name, world_rank, birth_year, federation, sex, fide_id, fide_title, standard_rating, rapid_ratig, blitz_rating)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
  RETURNING id, created_at, version
  `

	args := []any{p.Name, p.WorldRank, p.BYear, p.Federation, p.Sex, p.FideID, p.FideTttle, p.StandardRating, p.RapidRating, p.BlitzRating}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return pm.DB.QueryRow(ctx, query, args...).Scan(&p.ID, &p.CreatedAt, &p.Version)
}
