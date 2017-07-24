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
}

type theirs struct {
	Int    int
	String string
	fish   string
	Cat    string
}

func TestConfix(t *testing.T) {
	your := yours{
		An:          1,
		a:           "An A",
		TheirInt:    5,
		TheirString: "their string",
		Theirfish:   "their fish",
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
}
