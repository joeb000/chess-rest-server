# chess-rest-server - WORK IN PROGRESS
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
