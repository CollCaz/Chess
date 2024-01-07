CREATE TABLE IF NOT EXISTS games (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  version integer not null DEFAULT 1,
  event text,
  site text not null,
  date DATE not null,
  round integer not null,
  white text not null,
  black text not null,
  result text,
  pgn text
);
