package sportsndx

import (
	"testing"
)


func init() {
  initKeywords()
} // init


func TestCheckPlayerUniqueFirst(t *testing.T) {

	c := CheckPlayer("lebron")

	if c == nil || c[0].ID != 2544 {
		t.Error("Did not correctly classify player")
	}

} // TestCheckPlayerUniqueFirst


func TestCheckPlayerNonPlayer(t *testing.T) {

	c := CheckPlayer("leb")

	if c != nil {
		t.Error("Did not correctly classify player")
	}

} // TestCheckPlayerNonPlayer


func TestCheckPlayerUniqueLast(t *testing.T) {

	c := CheckPlayer("doncic")

	if c == nil || c[0].ID != 1629029 {
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


func TestSemanticParserSingleWord(t *testing.T) {

	c := SemanticParser("lebron")

	if len(c) != 1 {
		t.Error("Incorrect number of classifications")
	}

	if len(c) == 1 && !c[0].IsPlayer {
		t.Error("Did not correctly classify as player")
	}

} // TestSemanticParserSingleWord


func TestSemanticParserFirstLast(t *testing.T) {

	c := SemanticParser("lebron james")

	if len(c) != 2 {
		t.Error("Incorrect number of classifications")
	}

	if len(c) == 2 && !c[0].IsPlayer || !c[1].IsPlayer {
		t.Error("Did not correctly classify as player")
	}

} // TestSemanticParserFirstLast


func TestSemanticParserLastFirst(t *testing.T) {

	c := SemanticParser("james lebron")

	if len(c) != 2 {
		t.Error("Incorrect number of classifications")
	}

	if len(c) == 2 && !c[0].IsPlayer || !c[1].IsPlayer {
		t.Error("Did not correctly classify as player")
	}

} // TestSemanticParserLastFirst


func TestSemanticParserLast5(t *testing.T) {

	c := SemanticParser("lebron james last 5")

	if len(c) != 4 {
		t.Error("Incorrect number of classifications")
	}

	if len(c) == 4 && (!c[0].IsPlayer || !c[1].IsPlayer) && 
	  !c[2].IsKeyword && !c[3].IsParam {
		t.Error("Did not correctly classify")
	}

} // TestSemanticParserLast5


func TestSemanticParserEmptyString(t *testing.T) {

	c := SemanticParser("")

	if len(c) != 0 {
		t.Error("Incorrect number of classifications")
	}	

} // TestSemanticParserEmptyString


func TestSemanticParserCaseSensitivity(t *testing.T) {

	initKeywords()

	c := SemanticParser("lebron James")

	if c == nil {
		t.Error("Returned nil")
	}

} // TestSemanticParserCaseSensitivity


func TestSemanticParserTeamFull(t *testing.T) {

	initKeywords()

	c := SemanticParser("boston celtics")

	if len(c) != 2 {
		t.Error("Returned incorrect number of classifications")
	}

	if len(c) > 0 && !c[0].IsTeam {
		t.Error("Did not correctly classify as team")
	}

} // TestSemanticParserTeamFull


func TestSemanticParserTeamCity(t *testing.T) {

	initKeywords()

	c := SemanticParser("boston")

	if len(c) != 1 {
		t.Error("Returned incorrect classification number")
	}

	if len(c) == 1 && !c[0].IsTeam {
		t.Error("Did not correctly classify as team")
	}

} // TestSemanticParserTeamCity


func TestSemanticParserTeamName(t *testing.T) {

	initKeywords()

	c := SemanticParser("celtics")

	if len(c) != 1 {
		t.Error("Returned incorrect classification number")
	}

	if len(c) == 1 && !c[0].IsTeam {
		t.Error("Did not correctly classify as team")
	}

} // TestSemanticParserTeamName


func TestSemanticParserTeamAbv(t *testing.T) {

	initKeywords()

	c := SemanticParser("bos")

	if len(c) != 1 {
		t.Error("Returned incorrect classification number")
	}

	if len(c) == 1 && !c[0].IsTeam {
		t.Error("Did not correctly classify as team")
	}

} // TestSemanticParserTeamAbv


func TestSemanticParserTeamBadAbv(t *testing.T) {

	initKeywords()

	c := SemanticParser("bo")

	if len(c) != 1 {
		t.Error("Returned incorrect classification number")
	}

	if len(c) == 1 && c[0].IsTeam {
		t.Error("Classified incorrectly as team")
	}

} // TestSemanticParserTeamBadAbv
