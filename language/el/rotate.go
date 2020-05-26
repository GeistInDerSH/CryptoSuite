package el

import (
	"unicode"
)

// RotateN rotates the given string by n number of positions and returns the resulting string
func (el *EL) RotateN(s string, n int) (string, error) {
	var str string
	s = toUpper(s)

	isValid, err := el.isValidString(s)
	if !isValid {
		return "", err
	}

	for _, c := range []rune(s) {
		if unicode.IsSpace(c) || unicode.IsSymbol(c) || unicode.IsPunct(c) {
			str += string(c)
			continue
		}

		newVal := (el.runeAlphabet[c] + n) % el.alphabetLength
		str += string(el.intAlphabet[newVal])
	}
	return str, nil
}
