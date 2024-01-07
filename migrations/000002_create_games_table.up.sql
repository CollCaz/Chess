CREATE TABLE IF NOT EXISTS games (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  event text,
  site text not null,
  date DATE not null,
  round integer not null,
  white_id INT,
  black_id INT,
  result text,
  pgn text,
  
  CONSTRAINT fk_white
    FOREIGN KEY (white_id)
      REFERENCES players(id),
  
  CONSTRAINT fk_black
    FOREIGN KEY (black_id)
      REFERENCES players(id)
);
