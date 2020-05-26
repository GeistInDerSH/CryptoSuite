package en

import (
	"fmt"
	"strings"
	"unicode"
)

type EN struct {
	alphabetLength int
	runeAlphabet   map[rune]int
	intAlphabet    map[int]rune
	decodeMap      map[int]string
	encodeMap      map[string]string
}

// Init initializes the variables for the given EN variable
func (en *EN) Init() {
	en.decodeMap = map[int]string{
		11: "A", 12: "B", 13: "C", 14: "D", 15: "E",
		21: "F", 22: "G", 23: "H", 24: "I/J", 25: "K",
		31: "L", 32: "M", 33: "N", 34: "O", 35: "P",
		41: "Q", 42: "R", 43: "S", 44: "T", 45: "U",
		51: "V", 52: "W", 53: "X", 54: "Y", 55: "Z"}

	en.encodeMap = map[string]string{
		"A": "11", "B": "12", "C": "13", "D": "14", "E": "15",
		"F": "21", "G": "22", "H": "23", "I": "24", "J": "24", "K": "25",
		"L": "31", "M": "32", "N": "33", "O": "34", "P": "35",
		"Q": "41", "R": "42", "S": "43", "T": "44", "U": "45",
		"V": "51", "W": "52", "X": "53", "Y": "54", "Z": "55"}

	en.intAlphabet = map[int]rune{
		0: 'A', 1: 'B', 2: 'C', 3: 'D', 4: 'E',
		5: 'F', 6: 'G', 7: 'H', 8: 'I', 9: 'J',
		10: 'K', 11: 'L', 12: 'M', 13: 'N', 14: 'O',
		15: 'P', 16: 'Q', 17: 'R', 18: 'S', 19: 'T',
		20: 'U', 21: 'V', 22: 'W', 23: 'X', 24: 'Y',
		25: 'Z'}

	en.runeAlphabet = map[rune]int{
		'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4,
		'F': 5, 'G': 6, 'H': 7, 'I': 8, 'J': 9,
		'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14,
		'P': 15, 'Q': 16, 'R': 17, 'S': 18, 'T': 19,
		'U': 20, 'V': 21, 'W': 22, 'X': 23, 'Y': 24,
		'Z': 25}

	en.alphabetLength = len(en.runeAlphabet)

}

// stripSpace removes all spaces from the given string
func stripSpace(s string) string {
	var returnString string
	for _, c := range []rune(s) {
		if unicode.IsSpace(c) || unicode.IsControl(c) {
			continue
		}
		returnString += string(c)
	}

	return returnString
}

// toUpper converts the given string to all uppercase characters
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// isValidString validates if given string is valid in English
func (en *EN) isValidString(s string) (bool, error) {
	for _, c := range []rune(s) {
		if unicode.IsSpace(c) || unicode.IsSymbol(c) || unicode.IsPunct(c) {
			continue
		}

		_, result := en.runeAlphabet[c]
		if !result {
			err := fmt.Errorf("The given string is not valid for the language: %s", s)
			return false, err
		}
	}
	return true, nil
}
