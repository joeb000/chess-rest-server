package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"ChessIndex",
		"GET",
		"/chess",
		ChessIndex,
	},
	Route{
		"ChessMove",
		"POST",
		"/move",
		ChessMove,
	},
	Route{
		"ChessState",
		"GET",
		"/chess/{gameid}",
		DisplayBoard,
	},
	Route{
		"ChessCreate",
		"POST",
		"/create",
		ChessCreate,
	},
	Route{
		"ChessJoin",
		"POST",
		"/join/{gameid}",
		ChessJoin,
	},
	Route{
		"ChessFind",
		"GET",
		"/find/{gameid}",
		ChessFind,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TEST",
		"GET",
		"/test",
		TemplateHandle,
	},
}
