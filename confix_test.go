package confix

import (
	"testing"
)

type yours struct {
	An          int
	a           string
	TheirInt    int
	TheirString string
	Theirfish   string
	TheirNotInt int
}

type theirs struct {
	Int    int
	String string
	fish   string
	Cat    string
	NotInt string
}

func TestConfix(t *testing.T) {
	your := yours{
		An:          1,
		a:           "An A",
		TheirInt:    5,
		TheirString: "their string",
		Theirfish:   "their fish",
		TheirNotInt: 1,
	}

	var their theirs
	Confix("Their", &your, &their)
	if their.Int != 5 {
		t.Errorf("integer copy failed")
	}
	if their.String != "their string" {
		t.Errorf("string copy failed")
	}
	if their.Cat != "" {
		t.Errorf("accidentally set a string")
	}
	if their.fish != "" {
		t.Errorf("Privacy violation of unexported struct member")
	}
	if their.NotInt != "" {
		t.Errorf("Type violation of exported struct member")
	}
}
