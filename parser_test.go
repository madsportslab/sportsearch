package sportsndx

import (
	"testing"
)


func TestSemanticParser(t *testing.T) {

	res := SemanticParser("lebron james last 5")

	if res == nil {
		t.Error("Returned nil")
	}

} // TestSemanticParser


func TestSemanticParserEmptyString(t *testing.T) {

	res := SemanticParser("")

	if res != nil {
		t.Error("Returned non-nil")
	}

} // TestSemanticParserEmptyString


func TestSemanticParserCaseSensitivity(t *testing.T) {

	res := SemanticParser("lebron James")

	if res == nil {
		t.Error("Returned nil")
	}

} // TestSemanticParserCaseSensitivity
