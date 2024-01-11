
# Project Chess

Chess database API to store chess games and query them for various information.
This project was started as proof of what i learned before applying to a paid internship, I worked on it quickly and several features are missing.

## Done
- [x] Post a player
- [x] Post a game
- [x] Get a player by ID
- [x] Get a game by ID
- [x] Get all games from an abitrary query
## Todo
- [ ] Get all players from an arbitrary query
- [ ] Validate Inputs
- [ ] Send custom errors to the user
- [ ] More tests
- [ ] Store the players as a PostgreSQL object and use that in the Games instead of the player's names

## End Points
### Get /player/:id
  Get a player by their ID
### Post /player
  Add a player by sending json
  Example TODO
### Get /game/:id
  Get a game by ID
## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
	
