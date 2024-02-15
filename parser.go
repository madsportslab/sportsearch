package sportsndx

import (
	//"log"
  "strings"
)


type Who struct {
  Target  			string				`json:"target"`
	TargetType    int      			`json:"targetType"`
}


type What struct {
	Statistics    map[string]bool	`json:"statistics"`
}


type When struct {
	Period				int						`json:"period"`
	BeginRange		string        `json:"beginRange"`
	EndRange			string        `json:"endRange"`
	LastN         int           `json:"lastN"`
	LastNType     int           `json:"lastNType"`
}


type Intent struct {
  Players					[]int
	Teams						[]int
	LastN           int
	Begin      			string
	End             string
	Fields          []string
}


type Classification struct {
	IsPlayer          bool
	IsTeam            bool
	IsKeyword         bool
	IsParam           bool
}


func CheckPlayer(f string) []PlayerNameNode {

	var ret []PlayerNameNode

	_, ok := idxPlayerFirst[f]

	if ok {
		
		node := findByName(f, false)

		ret = make([]PlayerNameNode, len(node))

		copy(ret, node)

	}

	_, ok = idxPlayerLast[f]

	if ok {
		
		node := findByName(f, true)

		ret = make([]PlayerNameNode, len(node) + len(ret))

		copy(ret, node)

	}

	return ret

} // CheckPlayer


func CheckTeam(f string) int {

	lf := strings.ToLower(f)

	t, ok := idxTeamFull[lf]

	if ok {
		return t
	}

	t, ok = idxTeamName[lf]

	if ok {
		return t
	}

	t, ok = idxTeamCity[lf]

	if ok {
		return t
	}

	t, ok = idxTeamAbv[lf]

	if ok {
		return t
	}

	return NOT_FOUND

} // CheckTeam


func CheckKeyword(f string) {

} // CheckKeyword


func SemanticParser(q string) []Classification {

	if len(q) == 0 || q == EMPTY_STRING {
		return nil
	}

  fields := strings.Fields(q)

	var ret []Classification
	
	for _, f := range fields {

		c := Classification{}

		lf := strings.ToLower(f)

		CheckKeyword(lf)

		players := CheckPlayer(lf)

		if len(players) > 0 {
			c.IsPlayer = true
		}
		
		teamId := CheckTeam(lf)
		
		if teamId != NOT_FOUND {
			c.IsTeam = true
		}

		ret = append(ret, c)

	}

	return ret

} // SemanticParser
