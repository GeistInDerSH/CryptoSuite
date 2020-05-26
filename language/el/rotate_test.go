package el

import (
	"testing"
)

func TestRotateN1(t *testing.T) {
	e := new(EL)
	e.Init()

	n := 13

	cypherText, _ := e.RotateN("ΚΛΕΙΔΙ", n)

	expectedText := "ΨΩΣΧΡΧ"
	if cypherText != expectedText {
		t.Errorf("Incorrect Rotate %d\n\t\tGot: %s\n\t\tExpected: %s\n", n, cypherText, expectedText)
	}
}

func TestRotateN2(t *testing.T) {
	e := new(EL)
	e.Init()

	n := 13

	cypherText, _ := e.RotateN("ΨΩΣ ΧΡΧ", e.alphabetLength-n)

	expectedText := "ΚΛΕ ΙΔΙ"
	if cypherText != expectedText {
		t.Errorf("Incorrect Rotate %d\n\t\tGot: %s\n\t\tExpected: %s\n", n, cypherText, expectedText)
	}
}

func TestRotateN3(t *testing.T) {
	e := new(EL)
	e.Init()

	n := 13

	originalText := "ΚΛΕΙΔΙ"
	cypherText, _ := e.RotateN(originalText, n)
	plainText, _ := e.RotateN(cypherText, e.alphabetLength-n)

	if originalText != plainText {
		t.Errorf("Incorrect Rotate %d\n\t\tGot: %s\n\t\tExpected: %s\n", n, plainText, originalText)
	}
}

func TestRotateN4(t *testing.T) {
	e := new(EL)
	e.Init()

	n := 13

	_, err := e.RotateN("Cannot be rotated", n)

	if err == nil {
		t.Errorf("Incorrect Rotate %d\n\t\tGot: %s\n\t\tExpected: nil\n", n, err)
	}
}
