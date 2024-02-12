package sportsndx

import (
	"testing"
)


func TestInitTeams(t *testing.T) {

	initTeamIndexes()

	_, ok := teamIdxFull["Boston Celtics"]

	if !ok {
		t.Error("Not found")
	}

} // TestInitTeams


func TestInitPlayers(t *testing.T) {

	initPlayerIndexes()

	_, ok := playerIdxFull["Lebron James"]

	if !ok {
		t.Error("Not found")
	}

} // TestInitPlayers
