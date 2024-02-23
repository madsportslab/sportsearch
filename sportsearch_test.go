package sportsearch

import (
	"strings"
	"testing"
)


func init() {
	initKeywords()
} // init


func TestInitTeams(t *testing.T) {

	_, ok := idxTeamFull[strings.ToLower("Boston Celtics")]

	if !ok {
		t.Error("Not found")
	}

} // TestInitTeams


func TestInitPlayers(t *testing.T) {

	_, ok := idxPlayerFull[strings.ToLower("Lebron James")]

	if !ok {
		t.Error("Not found")
	}

} // TestInitPlayers


func TestFindByNameLast(t *testing.T) {

	players := findByName("Doncic", true)

	t.Log(players)
	
	if len(players) == 0 {
		t.Error("Not found")
	}

} // TestFindByNameLast


func TestFindByNameLastLowercase(t *testing.T) {

	players := findByName("doncic", true)

	t.Log(players)

	if len(players) == 0 {
		t.Error("Not found")
	}

} // TestFindByNameLastLowercase
