package el

import (
	"fmt"
	"strings"
	"unicode"
)

type EL struct {
	alphabetLength int
	runeAlphabet   map[rune]int
	intAlphabet    map[int]rune
	decodeMap      map[int]string
	encodeMap      map[string]string
}

// Init initializes the variables for the given EL variable
func (el *EL) Init() {
	el.decodeMap = map[int]string{
		11: "Α", 12: "Β", 13: "Γ", 14: "Δ", 15: "Ε",
		21: "Ζ", 22: "Η", 23: "Θ", 24: "Ι", 25: "Κ",
		31: "Λ", 32: "Μ", 33: "Ν", 34: "Ξ", 35: "Ο",
		41: "Π", 42: "Ρ", 43: "Σ", 44: "Τ", 45: "Υ",
		51: "Φ", 52: "Χ", 53: "Ψ", 54: "Ω"}

	el.encodeMap = map[string]string{
		"Α": "11", "Β": "12", "Γ": "13", "Δ": "14", "Ε": "15",
		"Ζ": "21", "Η": "22", "Θ": "23", "Ι": "24", "Κ": "25",
		"Λ": "31", "Μ": "32", "Ν": "33", "Ξ": "34", "Ο": "35",
		"Π": "41", "Ρ": "42", "Σ": "43", "Τ": "44", "Υ": "45",
		"Φ": "51", "Χ": "52", "Ψ": "53", "Ω": "54"}

	el.intAlphabet = map[int]rune{
		0: 'Α', 1: 'Β', 2: 'Γ', 3: 'Δ', 4: 'Ε',
		5: 'Ζ', 6: 'Η', 7: 'Θ', 8: 'Ι', 9: 'Κ',
		10: 'Λ', 11: 'Μ', 12: 'Ν', 13: 'Ξ', 14: 'Ο',
		15: 'Π', 16: 'Ρ', 17: 'Σ', 18: 'Τ', 19: 'Υ',
		20: 'Φ', 21: 'Χ', 22: 'Ψ', 23: 'Ω'}

	el.runeAlphabet = map[rune]int{
		'Α': 0, 'Β': 1, 'Γ': 2, 'Δ': 3, 'Ε': 4,
		'Ζ': 5, 'Η': 6, 'Θ': 7, 'Ι': 8, 'Κ': 9,
		'Λ': 10, 'Μ': 11, 'Ν': 12, 'Ξ': 13, 'Ο': 14,
		'Π': 15, 'Ρ': 16, 'Σ': 17, 'Τ': 18, 'Υ': 19,
		'Φ': 20, 'Χ': 21, 'Ψ': 22, 'Ω': 23}

	el.alphabetLength = len(el.runeAlphabet)
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

// isValidString checks if the given string is valid in Greek
func (el *EL) isValidString(s string) (bool, error) {
	for _, c := range []rune(s) {
		if unicode.IsSpace(c) || unicode.IsSymbol(c) || unicode.IsPunct(c) {
			continue
		}

		_, result := el.runeAlphabet[c]
		if !result {
			err := fmt.Errorf("The given string is not valid for the language: %s", s)
			return false, err
		}
	}
	return true, nil
}
