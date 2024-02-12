package sportsndx

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
	Accuracy				float32				`json:"accuracy"`
}


func SemanticParser(q string) *Intent {

	if len(q) == 0 || q == EMPTY_STRING {
		return nil
	}

  fields := strings.Fields(q)

	log.Println(fields)

	return &Intent{}

} // SemanticParser
