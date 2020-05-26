package en

import (
	"testing"
)

func TestVigenereEncodeEN1(t *testing.T) {
	e := new(EN)
	e.Init()
	plainText, _ := e.VigenereEncode("Key", "This is a test")

	expectedText := "DLGC MQ K XCCX"
	if plainText != expectedText {
		t.Errorf("Incorrect Encode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestVigenereEncodeEN2(t *testing.T) {
	e := new(EN)
	e.Init()
	_, err := e.VigenereEncode("ö", "Totaly normal text")

	if err == nil {
		t.Errorf("Incorrect Encode Error\n\t\tExpected: nil\n\t\tGot: %s\n", err)
	}
}

func TestVigenereEncodeEN3(t *testing.T) {
	e := new(EN)
	e.Init()
	_, err := e.VigenereEncode("Key", "Wir haben ein großes Problem")

	if err == nil {
		t.Errorf("Incorrect Encode Error\n\t\tExpected: nil\n\t\tGot: %s\n", err)
	}
}

func TestVigenereDecodeEN1(t *testing.T) {
	e := new(EN)
	e.Init()
	plainText, _ := e.VigenereDecode("ThisKeyisNotSecure", "GLQLRIPQKGVXLIZN")

	expectedText := "NEITHERISTHETEXT"
	if plainText != expectedText {
		t.Errorf("Incorrect Decode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestVigenereDecodeEN2(t *testing.T) {
	e := new(EN)
	e.Init()
	_, err := e.VigenereDecode("qwertyuiop", "1337 T3xT")

	if err == nil {
		t.Errorf("Incorrect Decode Error\n\t\tExpected: nil\n\t\tGot: %s\n", err)
	}
}

func TestVigenereDecodeEN3(t *testing.T) {
	e := new(EN)
	e.Init()
	_, err := e.VigenereDecode("ä", "Test String")

	if err == nil {
		t.Errorf("Incorrect Decode Error\n\t\tExpected: nil\n\t\tGot: %s\n", err)
	}
}

func TestVigenereDecodeEN4(t *testing.T) {
	e := new(EN)
	e.Init()
	_, err := e.VigenereDecode("!K y", "Test String")

	if err != nil {
		t.Errorf("Incorrect Decode Error\n\t\tExpected: nil\n\t\tGot: %s\n", err)
	}
}
