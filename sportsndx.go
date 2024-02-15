package sportsndx

import (
	//"log"
	"strings"
)


const (
	APP_NAME            = "sportsndx"
	APP_VERSION					= "0.1"
)


const (
	EMPTY_STRING				= ""
)


const (
	NOT_FOUND						= -1
)


const (
	KEYWORDS_DIR       	= "keywords"
	PLAYER_FILE					= "players.json"
	TEAM_FILE						= "teams.json"
	RESERVED_FILE       = "reserved.json"
)


const (
	COMPARATORS					=	"comparators"
	MATH                = "math"
	FIELDS              = "fields"
)


const (
  WHO_TEAM						= 0
	WHO_PLAYER         	= 1
	WHO_EAST            = 2
	WHO_WEST            = 3
	WHO_NONE            = 4
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


type PlayerNameNode struct {
	Name					string
	ID						int
}


type ReservedTypes struct {
  Comparators    			[]string			`json:"comparators"`
	Math    						[]string			`json:"math"`
	Fields    					[]string			`json:"fields"`
}


type AllKeywords struct {
	Keywords						ReservedTypes		`json:"keywords"`
}


// TODO: players that change names like metta world peace


var (
	idxTeamFull						map[string]int
	idxTeamName						map[string]int
	idxTeamCity						map[string]int
	idxTeamAbv						map[string]int
	idxPlayerFull					map[string]int
	idxPlayerFirst				map[string][]PlayerNameNode
	idxPlayerLast					map[string][]PlayerNameNode
	idxPlayerShort				map[string]int
	idxPlayerNicknames		map[string]int
	idxKeyComparators   	map[string]int
	idxKeyMath						map[string]int
	idxKeyTimeRange				map[string]int
	idxKeyFields					map[string]int
)


func setTeamMap(teams AllTeams) {

	for _, t := range teams.Teams {

		idxTeamFull[strings.ToLower(t.Full)] 	= t.ID
		idxTeamName[strings.ToLower(t.Name)] 	= t.ID
		idxTeamCity[strings.ToLower(t.City)] 	= t.ID
		idxTeamAbv[strings.ToLower(t.Abv)] 		= t.ID

	}

} // setTeamMap


func setPlayerMap(players AllPlayers) {

	for _, p := range players.Players {

		last 	:= strings.ToLower(p.Last)
		first := strings.ToLower(p.First)
		full 	:= strings.ToLower(p.Full)
		short := strings.ToLower(p.Short)

		idxPlayerFull[full] 		= p.ID
		idxPlayerShort[short] 	= p.ID

		idxPlayerLast[last] =
		  append(idxPlayerLast[last],			
		  PlayerNameNode{
				Name: first,
				ID: p.ID,
			})     

		idxPlayerFirst[first] =
		  append(idxPlayerFirst[first],
			PlayerNameNode{
			  Name: last,
			  ID: p.ID,
			})

		for _, n := range p.Nicknames {
			idxPlayerNicknames[strings.ToLower(n)] = p.ID
		}

	}

} // setPlayerMap


func setKeywordMap(keywords AllKeywords) {

	for _, c := range keywords.Keywords.Comparators {
		idxKeyComparators[strings.ToLower(c)] = 1
	}

	for _, m := range keywords.Keywords.Math {
		idxKeyMath[strings.ToLower(m)] = 1
	}

  for _, f := range keywords.Keywords.Fields {
		idxKeyFields[strings.ToLower(f)] = 1
	}
	
} // setKeywordMap


func findByName(n string, isLast bool) []PlayerNameNode {

	var idx map[string][]PlayerNameNode

	if len(n) == 0 || n == EMPTY_STRING {
		return nil
	}

	if isLast {
		idx = idxPlayerLast
	} else {
		idx = idxPlayerFirst
	}	

	ln := strings.ToLower(n)
	
	return idx[ln]

} // findByName


func initPlayerIndexes() {

	idxPlayerFull 			= make(map[string]int)
	idxPlayerLast 			= make(map[string][]PlayerNameNode)
	idxPlayerFirst 			= make(map[string][]PlayerNameNode)
	idxPlayerShort 			= make(map[string]int)
	idxPlayerNicknames 	= make(map[string]int)

	players := AllPlayers{}

	ReadJson(PLAYER_FILE, &players)

  setPlayerMap(players)

} // initPlayerIndexes


func initTeamIndexes() {

	idxTeamFull 	= make(map[string]int)
	idxTeamName 	= make(map[string]int)
	idxTeamCity 	= make(map[string]int)
	idxTeamAbv 		= make(map[string]int)

	teams := AllTeams{}

	ReadJson(TEAM_FILE, &teams)

  setTeamMap(teams)

} // initTeamIndexes


func initReservedIndexes() {

	idxKeyComparators		= make(map[string]int)
	idxKeyMath					= make(map[string]int)
	idxKeyFields				= make(map[string]int)

	keywords := AllKeywords{}

	ReadJson(RESERVED_FILE, &keywords)

  setKeywordMap(keywords)

} // initReservedIndexes


func initKeywords() {

	initTeamIndexes()
	initPlayerIndexes()
	initReservedIndexes()

} // initKeywords
