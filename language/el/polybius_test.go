package el

import (
	"testing"
)

func TestPolybiusEncodeELCyphertext1(t *testing.T) {
	e := new(EL)
	e.Init()
	cypherText, _ := e.PolybiusEncode("ΚΛΕΙΔΙ")

	expectedText := "25 31 15 24 14 24 "
	if cypherText != expectedText {
		t.Errorf("Incorrect Encode\n\t\tGot: %s\n\t\tExpected: %s\n", cypherText, expectedText)
	}
}

func TestPolybiusEncodeELCyphertext2(t *testing.T) {
	e := new(EL)
	e.Init()
	cypherText, _ := e.PolybiusEncode("ΑΝΑΤΟΛΗ")

	expectedText := "11 33 11 44 35 31 22 "
	if cypherText != expectedText {
		t.Errorf("Incorrect Encode\n\t\tGot: %s\n\t\tExpected: %s\n", cypherText, expectedText)
	}
}

func TestPolybiusEncodeELError1(t *testing.T) {
	e := new(EL)
	e.Init()
	_, err := e.PolybiusEncode("This is a Test")

	if err == nil {
		t.Errorf("Incorrect Encode\n\t\tGot: %s\n\t\tExpected: nil\n", err)
	}
}

func TestPolybiusDecodeELCyphertext1(t *testing.T) {
	e := new(EL)
	e.Init()
	cypherText, _ := e.PolybiusDecode("25 31 15 24 14 24")

	expectedText := "ΚΛΕΙΔΙ"
	if cypherText != expectedText {
		t.Errorf("Incorrect Decode\n\t\tGot: %s\n\t\tExpected: %s\n", cypherText, expectedText)
	}
}

func TestPolybiusDecodeELCyphertext2(t *testing.T) {
	e := new(EL)
	e.Init()
	cypherText, _ := e.PolybiusDecode("11 33 11 44 35 31 22")

	expectedText := "ΑΝΑΤΟΛΗ"
	if cypherText != expectedText {
		t.Errorf("Incorrect Decode\n\t\tGot: %s\n\t\tExpected: %s\n", cypherText, expectedText)
	}
}

func TestPolybiusDecodeELError1(t *testing.T) {
	e := new(EL)
	e.Init()
	_, err := e.PolybiusDecode("25 31 15 24 14 24!")

	if err == nil {
		t.Errorf("Incorrect Decode\n\t\tGot: %s\n\t\tExpected: nil\n", err)
	}
}

func TestPolybiusDecodeELError2(t *testing.T) {
	e := new(EL)
	e.Init()
	_, err := e.PolybiusDecode("aa")

	if err == nil {
		t.Errorf("Incorrect Decode\n\t\tGot: %s\n\t\tExpected: nil\n", err)
	}
}
