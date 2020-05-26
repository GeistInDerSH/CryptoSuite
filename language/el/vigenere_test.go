package el

import (
	"testing"
)

func TestVigenereEncodeELPlaintext1(t *testing.T) {
	e := new(EL)
	e.Init()
	plainText, _ := e.VigenereEncode("ΥΠΟ", "ΒΟΡΕΙΟΣ")

	expectedText := "ΦΖΗΩΩΕΝ"
	if plainText != expectedText {
		t.Errorf("Incrrect Encode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestVigenereEncodeELPlaintext2(t *testing.T) {
	e := new(EL)
	e.Init()
	plainText, _ := e.VigenereEncode("ΥΠΟ", "ΚΛΕ ΙΔΙ")

	expectedText := "ΕΒΤ ΔΤΨ"
	if plainText != expectedText {
		t.Errorf("Incrrect Encode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestVigenereEncodeELError1(t *testing.T) {
	e := new(EL)
	e.Init()
	_, err := e.VigenereEncode("Key", "This is in English")

	if err == nil {
		t.Errorf("Incrrect Encode\n\t\tGot: %s\n\t\tExpected: nil\n", err)
	}
}

func TestVigenereEncodeELError2(t *testing.T) {
	e := new(EL)
	e.Init()
	_, err := e.VigenereEncode("NotValid", "ΕΒΤ ΔΤΨ")

	if err == nil {
		t.Errorf("Incrrect Encode\n\t\tGot: %s\n\t\tExpected: nil\n", err)
	}
}

func TestVigenereDecodeELPlaintext1(t *testing.T) {
	e := new(EL)
	e.Init()
	plainText, _ := e.VigenereDecode("ΥΠΟ", "ΦΖΗΩΩΕΝ")

	expectedText := "ΒΟΡΕΙΟΣ"
	if plainText != expectedText {
		t.Errorf("Incrrect Decode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestVigenereDecodeELPlaintext2(t *testing.T) {
	e := new(EL)
	e.Init()
	plainText, _ := e.VigenereDecode("ΥΠΟ", "ΕΒΤ ΔΤΨ")

	expectedText := "ΚΛΕ ΙΔΙ"
	if plainText != expectedText {
		t.Errorf("Incrrect Decode\n\t\tGot: %s\n\t\tExpected: %s\n", plainText, expectedText)
	}
}

func TestVigenereDecodeELError1(t *testing.T) {
	e := new(EL)
	e.Init()
	_, err := e.VigenereDecode("ΥΠΟ", "This is a test")

	if err == nil {
		t.Errorf("Incrrect Decode\n\t\tGot: %s\n\t\tExpected: nil\n", err)
	}
}

func TestVigenereDecodeELError2(t *testing.T) {
	e := new(EL)
	e.Init()
	_, err := e.VigenereDecode("NotValid", "ΕΒΤ ΔΤΨ")

	if err == nil {
		t.Errorf("Incrrect Decode\n\t\tGot: %s\n\t\tExpected: nil\n", err)
	}
}
