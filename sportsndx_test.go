package sportsndx

import (
	"strings"
	"testing"
)


func TestInitTeams(t *testing.T) {

	initTeamIndexes()

	_, ok := idxTeamFull[strings.ToLower("Boston Celtics")]

	if !ok {
		t.Error("Not found")
	}

} // TestInitTeams


func TestInitPlayers(t *testing.T) {

	initPlayerIndexes()

	_, ok := idxPlayerFull[strings.ToLower("Lebron James")]

	if !ok {
		t.Error("Not found")
	}

} // TestInitPlayers
