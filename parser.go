package sportsearch

import (
	"log"
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
	Players           []int
	Teams             []int
	Fields            []string
	Math              []string
}


type PageSummary struct {
  
}

func CheckPlayer(f string, c *Classification, isLast bool) {

	var idx map[string][]PlayerNameNode

	if c.Players == nil {
		c.Players = []int{}
	}

	if isLast {
		idx = idxPlayerLast
	} else {
		idx = idxPlayerFirst
	}

	fl := strings.ToLower(f)

	_, ok := idx[fl]

	if ok {
		
		c.IsPlayer = true

		nodes := findByName(fl, isLast)

		for _, n := range nodes {
			c.Players = append(c.Players, n.ID)
		}

	}

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


func CheckKeyword(f string) int {

	strings.ToLower(f)

	return NOT_FOUND

} // CheckKeyword


func Classifier(q string) []Classification {

	if len(q) == 0 || q == EMPTY_STRING {
		return nil
	}

  fields := strings.Fields(q)

	var ret []Classification
	
	for _, f := range fields {

		c := Classification{}

		lf := strings.ToLower(f)

		CheckPlayer(lf, &c, true)
		CheckPlayer(lf, &c, false)
		
		teamId := CheckTeam(lf)
		
		if teamId != NOT_FOUND {
			c.IsTeam = true
		}

		ret = append(ret, c)

		key := CheckKeyword(lf)

		if key != NOT_FOUND {
			c.IsKeyword = true
		}

	}

	log.Printf("%v\n", ret)

	return ret

} // Classifier


func search() []PageSummary {
  return nil
} // search
