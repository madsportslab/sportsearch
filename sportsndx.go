package sportsndx

import (
	//"log"
)


const (
	APP_NAME            = "Sportsndx"
	APP_VERSION					= "0.1"
)


const (
	EMPTY_STRING				= ""
)


const (
	SEMANTICS_DIR       = "semantics"
	PLAYER_FILE					= "players.json"
	TEAM_FILE						= "teams.json"
)


type Team struct {
  Full					string 			`json:"full"`
	Name					string 			`json:"name"`
	City					string 			`json:"city"`
	Abv						string 			`json:"abv"`
	ID						int					`json:"id"`
}

type AllTeams struct {
	Teams 				[]Team			`json:"teams"`
}


type Player struct {
  First					string 			`json:"first"`
	Last					string 			`json:"last"`
	Full					string 			`json:"full"`
	Short					string 			`json:"short"`
	ID						int 				`json:"id"`
	Nicknames			[]string 		`json:"nicknames"`
}

type AllPlayers struct {
	Players      []Player   `json:"players"`
}

var (
	teamIdxFull						map[string]int
	teamIdxName						map[string]int
	teamIdxCity						map[string]int
	teamIdxAbv						map[string]int
	playerIdxFull					map[string]int
	//playerIdxFirst				map[string]int
	//playerIdxLast					map[string]int
	playerIdxShort				map[string]int
	playerIdxNicknames		map[string]int
)


func setTeamMap(teams AllTeams) {

	for _, t := range teams.Teams {

		teamIdxFull[t.Full] 	= t.ID
		teamIdxName[t.Name] 	= t.ID
		teamIdxCity[t.City] 	= t.ID
		teamIdxAbv[t.Abv] 		= t.ID

	}

} // setTeamMap


func setPlayerMap(players AllPlayers) {

	for _, p := range players.Players {

		playerIdxFull[p.Full] 	= p.ID
		playerIdxShort[p.Short] 	= p.ID

		for _, n := range p.Nicknames {
			playerIdxNicknames[n] = p.ID
		}

	}

} // setPlayerMap


func initPlayerIndexes() {

	playerIdxFull 			= make(map[string]int)
	//playerIdxLast 			= make(map[string]int)
	playerIdxShort 			= make(map[string]int)
	playerIdxNicknames 	= make(map[string]int)

	players := AllPlayers{}

	ReadJson(PLAYER_FILE, &players)

  setPlayerMap(players)

} // initPlayerIndexes


func initTeamIndexes() {

	teamIdxFull 	= make(map[string]int)
	teamIdxName 	= make(map[string]int)
	teamIdxCity 	= make(map[string]int)
	teamIdxAbv 		= make(map[string]int)

	teams := AllTeams{}

	ReadJson(TEAM_FILE, &teams)

  setTeamMap(teams)

} // initTeamIndexes


func initKeywords() {

	initTeamIndexes()
	initPlayerIndexes()

} // initKeywords
