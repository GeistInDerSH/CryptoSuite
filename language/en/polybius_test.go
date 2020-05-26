package en

import (
	"testing"
)

func TestPolybiusDecodeEN1(t *testing.T) {
	e := new(EN)
	e.Init()
	plainText, _ := e.PolybiusDecode("44 23 15 41 45 24 13 25 12 42 34 52 33 21 34 53 45 32 35 43 34 51 15 42 44 23 15 31 11 55 54 14 34 22")

	expectedText := "THEQUI/JCKBROWNFOXUMPSOVERTHELAZYDOG"
	if plainText != expectedText {
		t.Errorf("Incorrect Decode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestPolybiusDecodeEN2(t *testing.T) {
	e := new(EN)
	e.Init()
	plainText, _ := e.PolybiusDecode("221524434424331415424323")

	expectedText := "GEI/JSTI/JNDERSH"
	if plainText != expectedText {
		t.Errorf("Incorrect Decode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestPolybiusDecodeEN3(t *testing.T) {
	e := new(EN)
	e.Init()
	_, err := e.PolybiusDecode("1121!!")

	if err == nil {
		t.Errorf("Incorrect Decode Error\n\t\tExpected: nil\n\t\tGot: %s\n", err)
	}
}

func TestPolybiusDecodeEN4(t *testing.T) {
	e := new(EN)
	e.Init()
	_, err := e.PolybiusDecode("This is a test")

	if err == nil {
		t.Errorf("Incorrect Decode Error\n\t\tExpected: nil\n\t\tGot: %s\n", err)
	}
}

func TestPolybiusEncodeEN1(t *testing.T) {
	e := new(EN)
	e.Init()
	plainText, _ := e.PolybiusEncode("the quick brown fox jumps over the lazy dog")

	expectedText := "44 23 15 41 45 24 13 25 12 42 34 52 33 21 34 53 24 45 32 35 43 34 51 15 42 44 23 15 31 11 55 54 14 34 22 "
	if plainText != expectedText {
		t.Errorf("Incorrect Decode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestPolybiusEncodeEN2(t *testing.T) {
	e := new(EN)
	e.Init()
	plainText, _ := e.PolybiusEncode("j")

	expectedText := "24 "
	if plainText != expectedText {
		t.Errorf("Incorrect Encode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestPolybiusEncodeEN3(t *testing.T) {
	e := new(EN)
	e.Init()
	_, err := e.PolybiusEncode("1121!!")

	if err == nil {
		t.Errorf("Incorrect Encode Error\n\t\tExpected: nil\n\t\tGot: %s\n", err)
	}
}

func TestPolybiusEncodeEN4(t *testing.T) {
	e := new(EN)
	e.Init()
	_, err := e.PolybiusEncode("Th1$ h4$ numb3r$!")

	if err == nil {
		t.Errorf("Incorrect Encode Error\n\t\tExpected: nil\n\t\tGot: %s\n", err)
	}
}
