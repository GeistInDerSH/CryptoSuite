package el

import (
	"fmt"
	"strconv"
)

// PolybiusDecode attempts to decode a string of numbers into Greek characters.
// If there are any errors the function will return an empty string
// and the error
func (el EL) PolybiusDecode(s string) (result string, err error) {

	if len(stripSpace(s))%2 != 0 {
		err := fmt.Errorf("The string to be decoded is not valid")
		return "", err

	}
	s = stripSpace(s)

	for i := 0; i < len(s); i += 2 {
		integer, err := strconv.Atoi(s[i : i+2])
		if err != nil {
			return "", fmt.Errorf("Could not convert \"%s\" to integer", s[i:i+2])
		}
		result += el.decodeMap[integer]
	}
	return
}

// PolybiusEncode encodes the given string into the  polybius number values
func (el EL) PolybiusEncode(s string) (string, error) {
	var result string
	s = toUpper(stripSpace(s))

	isValidString, err := el.isValidString(s)
	if !isValidString {
		return "", err
	}

	for _, i := range []rune(s) {
		result += el.encodeMap[string(i)] + " "
	}
	return result, nil
}
