# ♜  chess-rest-server  ♜
RESTful API for managing chess gameplay written in golang

## Clone and install

    git clone https://github.com/joeb000/chess-rest-server
    cd chess-rest-server/main
    go get
    go run *.go
    
## Using the API

Chess server exposes HTTP endpoints that can handle `GET` and `POST` calls.

    GET:
    /chess - returns welcome message
    
    /chess/{gameid} - returns JSON Game object for a given gameid
    
    

    POST:
    /create - takes a JSON Player object and initializes an empty game, returns Game object
    
    /join/{gameid} - takes a JSON Player object and joins empty game at given gameid, returns Game object
    
    /move - takes a Move JSON object and updates the game state accordingly - returns game object


## Example Gameplay
    
    # To Create a new game as player "Alice"
    curl -H "Content-Type: application/json" -d '{"Id":10, "Name":"Alice"}' http://localhost:8080/create
    
    # To Join Game ID 1 as player "Bob"
    curl -H "Content-Type: application/json" -d '{"Id":11, "Name":"Bob"}' http://localhost:8080/join/1
    
    # To Move piece from square A1 to square A2
    curl -H "Content-Type: application/json" -d '{"game_id":1, "player_id":10, "from_square":"A1", "to_square":"A2"}' http://localhost:8080/move
    
To check game state: http://localhost:8080/find/1

