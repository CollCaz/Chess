package data

import (
	"context"
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PlayerModel struct {
	DB *pgxpool.Pool
}

type Player struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdat"`
	WorldRank      int       `json:"worldrank"`
	BYear          int       `json:"byear"`
	Federation     string    `json:"federation"`
	Sex            string    `json:"sex"`
	FideID         int       `json:"fideid"`
	FideTttle      string    `json:"fidetitle"`
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
		FirstName  string
		LastName   string
		WorldRank  int
		BYear      int
		Federation string
		Sex        string
		FideID     int
		FideTttle  string
		Ratings    map[string]int
	}{
		Name:       p.Name,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		WorldRank:  p.WorldRank,
		BYear:      p.BYear,
		Federation: p.Federation,
		Sex:        p.Sex,
		FideID:     p.FideID,
		FideTttle:  p.FideTttle,
		Ratings:    rating,
	}

	return json.Marshal(aux)
}

func (pm PlayerModel) GetPlayer(id int) (*Player, error) {
	var player Player
	/*
		magnus := Player{
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
	*/

	query := `
  SELECT id, name, first_name, last_name, version, created_at, world_rank, birth_year, federation, sex, fide_id, fide_title, standard_rating, rapid_rating, blitz_rating
  FROM players
  WHERE id = $1
  `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := pm.DB.QueryRow(ctx, query, id).Scan(
		&player.ID,
		&player.Name,
		&player.FirstName,
		&player.LastName,
		&player.Version,
		&player.CreatedAt,
		&player.WorldRank,
		&player.BYear,
		&player.Federation,
		&player.Sex,
		&player.FideID,
		&player.FideTttle,
		&player.StandardRating,
		&player.RapidRating,
		&player.BlitzRating,
	)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (pm PlayerModel) InsertPlayer(p *Player) error {
	query := `
  INSERT INTO players (name, first_name, last_name, world_rank, birth_year, federation, sex, fide_id, fide_title, standard_rating, rapid_rating, blitz_rating)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)

  RETURNING id, created_at, version
  `

	args := []any{&p.Name, &p.FirstName, &p.LastName, &p.WorldRank, &p.BYear, &p.Federation, &p.Sex, &p.FideID, &p.FideTttle, &p.StandardRating, &p.RapidRating, &p.BlitzRating}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return pm.DB.QueryRow(ctx, query, args...).Scan(&p.ID, &p.CreatedAt, &p.Version)
}
