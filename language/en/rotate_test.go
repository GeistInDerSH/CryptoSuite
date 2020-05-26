package en

import (
	"testing"
)

func TestRotateN1(t *testing.T) {
	e := new(EN)
	e.Init()

	n := 13

	cypherText, _ := e.RotateN("Ceasar Cypher", n)

	expectedText := "PRNFNE PLCURE"
	if cypherText != expectedText {
		t.Errorf("Incorrect Rotate %d\n\t\tGot: %s\n\t\tExpected: %s\n", n, cypherText, expectedText)
	}
}

func TestRotateN2(t *testing.T) {
	e := new(EN)
	e.Init()

	n := 13

	cypherText, _ := e.RotateN("Hello World!", n)

	expectedText := "URYYB JBEYQ!"
	if cypherText != expectedText {
		t.Errorf("Incorrect Rotate %d\n\t\tGot: %s\n\t\tExpected: %s\n", n, cypherText, expectedText)
	}
}

func TestRotateN3(t *testing.T) {
	e := new(EN)
	e.Init()

	n := 0

	cypherText, _ := e.RotateN("They are exactly the same", n)

	expectedText := "THEY ARE EXACTLY THE SAME"
	if cypherText != expectedText {
		t.Errorf("Incorrect Rotate %d\n\t\tGot: %s\n\t\tExpected: %s\n", n, cypherText, expectedText)
	}
}

func TestRotateN4(t *testing.T) {
	e := new(EN)
	e.Init()

	n := 7

	originalText := "IPSUM DOLOR SIT AMET"
	cypherText, _ := e.RotateN(originalText, n)
	plainText, _ := e.RotateN(cypherText, e.alphabetLength-n)

	if plainText != originalText {
		t.Errorf("Incorrect Rotate %d\n\t\tGot: %s\n\t\tExpected: %s\n", n, plainText, originalText)
	}
}

func TestRotateN5(t *testing.T) {
	e := new(EN)
	e.Init()

	n := 13

	_, err := e.RotateN("Cäsar Cypher", n)

	if err == nil {
		t.Errorf("Incorrect Rotate %d\n\t\tExpected: nil\n\t\tGot: %s\n", n, err)
	}
}
