package sportsndx

import (
	"testing"
)


func init() {
  initKeywords()
} // init


func TestCheckPlayerUniqueFirst(t *testing.T) {

	c := Classification{}

	CheckPlayer("lebron", &c, true)
	CheckPlayer("lebron", &c, false)

	if !c.IsPlayer || c.Players[0] != 2544 {
		t.Error("Did not correctly classify player")
	}

} // TestCheckPlayerUniqueFirst


func TestCheckPlayerNonPlayer(t *testing.T) {

	c := Classification{}

	CheckPlayer("leb", &c, true)
	CheckPlayer("leb", &c, false)

	if !c.IsPlayer || c.Players[0] != 2544 {
		t.Error("Did not correctly classify player")
	}

} // TestCheckPlayerNonPlayer


func TestCheckPlayerUniqueLast(t *testing.T) {

	c := Classification{}

	CheckPlayer("doncic", &c, true)
	CheckPlayer("doncic", &c, false)

	if !c.IsPlayer || c.Players[0] != 1629029 {
		t.Error("Did not correctly classify player")
	}

} // TestCheckPlayerUniqueLast


func TestCheckTeamCity(t *testing.T) {

	c := CheckTeam("boston")

	if c == NOT_FOUND || c != 1610612738 {
		t.Error("Did not correctly classify team")
	}

} // TestCheckTeamCity


func TestCheckTeamName(t *testing.T) {

	c := CheckTeam("celtics")

	if c == NOT_FOUND || c != 1610612738 {
		t.Error("Did not correctly classify team")
	}

} // TestCheckTeamName


func TestCheckTeamFull(t *testing.T) {

	c := CheckTeam("boston celtics")

	if c == NOT_FOUND || c != 1610612738 {
		t.Error("Did not correctly classify team")
	}

} // TestCheckTeamFull


func TestCheckTeamAbv(t *testing.T) {

	c := CheckTeam("bos")

	if c == NOT_FOUND || c != 1610612738 {
		t.Error("Did not correctly classify team")
	}

} // TestCheckTeamAbv


func TestCheckTeamNotExist(t *testing.T) {

	c := CheckTeam("shanghai")

	if c != NOT_FOUND {
		t.Error("Team does not exist, should not have a match")
	}

} // TestCheckTeamNotExist


func TestClassifierSingleWord(t *testing.T) {

	c := Classifier("lebron")

	if len(c) != 1 {
		t.Error("Incorrect number of classifications")
	}

	if len(c) == 1 && !c[0].IsPlayer {
		t.Error("Did not correctly classify as player")
	}

} // TestClassifierSingleWord


func TestClassifierFirstLast(t *testing.T) {

	c := Classifier("lebron james")

	if len(c) != 2 {
		t.Error("Incorrect number of classifications")
	}

	if len(c) == 2 && !c[0].IsPlayer || !c[1].IsPlayer {
		t.Error("Did not correctly classify as player")
	}

} // TestClassifierFirstLast


func TestClassifierLastFirst(t *testing.T) {

	c := Classifier("james lebron")

	if len(c) != 2 {
		t.Error("Incorrect number of classifications")
	}

	if len(c) == 2 && !c[0].IsPlayer || !c[1].IsPlayer {
		t.Error("Did not correctly classify as player")
	}

} // TestClassifierLastFirst


func TestClassifierLast5(t *testing.T) {

	c := Classifier("lebron james last 5")

	if len(c) != 4 {
		t.Error("Incorrect number of classifications")
	}

	if len(c) == 4 && (!c[0].IsPlayer || !c[1].IsPlayer) && 
	  !c[2].IsKeyword && !c[3].IsParam {
		t.Error("Did not correctly classify")
	}

} // TestClassifierLast5


func TestClassifierEmptyString(t *testing.T) {

	c := Classifier("")

	if len(c) != 0 {
		t.Error("Incorrect number of classifications")
	}	

} // TestClassifierEmptyString


func TestClassifierCaseSensitivity(t *testing.T) {

	initKeywords()

	c := Classifier("lebron James")

	if c == nil {
		t.Error("Returned nil")
	}

} // TestClassifierCaseSensitivity


func TestClassifierTeamFull(t *testing.T) {

	initKeywords()

	c := Classifier("boston celtics")

	if len(c) != 2 {
		t.Error("Returned incorrect number of classifications")
	}

	if len(c) > 0 && !c[0].IsTeam {
		t.Error("Did not correctly classify as team")
	}

} // TestClassifierTeamFull


func TestClassifierTeamCity(t *testing.T) {

	initKeywords()

	c := Classifier("boston")

	if len(c) != 1 {
		t.Error("Returned incorrect classification number")
	}

	if len(c) == 1 && !c[0].IsTeam {
		t.Error("Did not correctly classify as team")
	}

} // TestClassifierTeamCity


func TestClassifierTeamName(t *testing.T) {

	initKeywords()

	c := Classifier("celtics")

	if len(c) != 1 {
		t.Error("Returned incorrect classification number")
	}

	if len(c) == 1 && !c[0].IsTeam {
		t.Error("Did not correctly classify as team")
	}

} // TestClassifierTeamName


func TestClassifierTeamAbv(t *testing.T) {

	initKeywords()

	c := Classifier("bos")

	if len(c) != 1 {
		t.Error("Returned incorrect classification number")
	}

	if len(c) == 1 && !c[0].IsTeam {
		t.Error("Did not correctly classify as team")
	}

} // TestClassifierTeamAbv


func TestClassifierTeamBadAbv(t *testing.T) {

	initKeywords()

	c := Classifier("bo")

	if len(c) != 1 {
		t.Error("Returned incorrect classification number")
	}

	if len(c) == 1 && c[0].IsTeam {
		t.Error("Classified incorrectly as team")
	}

} // TestClassifierTeamBadAbv
