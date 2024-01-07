CREATE TABLE IF NOT EXISTS players (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  name text NOT NULL,
  world_rank integer NOT NULL,
  birth_year integer NOT NULL,
  federation text NOT NULL,
  sex text NOT NULL,
  fide_id integer UNIQUE NOT NULL,
  fide_title text NOT NULL,
  standard_rating integer NOT NULL,
  rapid_rating integer NOT NULL,
  blitz_rating integer NOT NULL,
  version integer NOT NULL DEFAULT 1
);
