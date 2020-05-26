package en

import (
	"unicode"
)

// extendKey extends the given key untill it matches the length of the plaintext.
// Any unicode spaces or symbols in the plaintext are skipped.
// E.g. key = "KEY", text = "THIS IS A TEST" results in: "KEYK EY K EYKE"
func extendKey(key, text string) []rune {
	var extendedKey string
	var i int = -1
	runeKey := []rune(key)

	for _, c := range []rune(text) {
		i++
		if unicode.IsSpace(c) || unicode.IsSymbol(c) {
			extendedKey += string(c)
			i--
		} else {
			if i == len(key) {
				i = 0
			}
			extendedKey += string(runeKey[i])
		}
	}

	return []rune(extendedKey)
}

// VigenereEncode encodes the plaintext with the key
// If either the key or plaintext are not valid for the language, an error with be returned
func (en *EN) VigenereEncode(key, text string) (string, error) {
	key = toUpper(key)
	text = toUpper(text)
	isValid, err := en.isValidString(text)
	if !isValid {
		return "", err
	}
	isValid, err = en.isValidString(key)
	if !isValid {
		return "", err
	}

	extendedKey := extendKey(key, text)
	var cyphertext string
	for i, c := range []rune(text) {
		if unicode.IsSpace(c) || unicode.IsSymbol(c) {
			cyphertext += string(c)
			continue
		}
		newPosition := (en.runeAlphabet[c] + en.runeAlphabet[extendedKey[i]]) % en.alphabetLength
		cyphertext += string(en.intAlphabet[newPosition])
	}

	return cyphertext, nil
}

// VigenereDecode decodes the cyphertext with the key
// If either the key or cyphertext are not valid for the language, an error with be returned
func (en *EN) VigenereDecode(key, text string) (string, error) {
	key = toUpper(key)
	text = toUpper(text)

	isValid, err := en.isValidString(text)
	if !isValid {
		return "", err
	}
	isValid, err = en.isValidString(key)
	if !isValid {
		return "", err
	}

	extendedKey := extendKey(key, text)

	var plaintext string
	for i, c := range []rune(text) {
		if unicode.IsSpace(c) || unicode.IsSymbol(c) {
			plaintext += string(c)
			continue
		}
		newPosition := (en.runeAlphabet[c] - en.runeAlphabet[extendedKey[i]] + en.alphabetLength) % en.alphabetLength
		plaintext += string(en.intAlphabet[newPosition])
	}

	return plaintext, nil
}
