package en

import (
	"unicode"
)

// RotateN rotates the given string by n number of positions and returns the resulting string
func (en *EN) RotateN(s string, n int) (string, error) {
	var str string
	s = toUpper(s)

	isValid, err := en.isValidString(s)
	if !isValid {
		return "", err
	}

	for _, c := range []rune(s) {
		if unicode.IsSpace(c) || unicode.IsSymbol(c) || unicode.IsPunct(c) {
			str += string(c)
			continue
		}

		newVal := (en.runeAlphabet[c] + n) % en.alphabetLength
		str += string(en.intAlphabet[newVal])
	}
	return str, nil
}
