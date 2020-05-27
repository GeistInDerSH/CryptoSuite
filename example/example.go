package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	crypto "github.com/GeistInDerSH/CryptoSuite"
)

var (
	lang     = flag.String("l", "", "String for what language to use\nThis flag is required\nE.g. en (English) or el (Greek)")
	fileName = flag.String("f", "", "Read from a file instead of stdin\nUsage: -f <file path>")

	isRotateN    = flag.Bool("rot", false, "Use a rotate N cypher")
	rotateAmount = flag.Int("n", 0, "The number of positions to rotate each character by")

	isVigenere       = flag.Bool("ver", false, "Use a Vigenere Cypher")
	vigenereKey      = flag.String("k", "", "Encryption/Decryption key for a Vigenere Cypher")
	isPolybiusSquare = flag.Bool("ps", false, "Enable Polybius Square")
	encodeBool       = flag.Bool("e", false, "Encode a given string")
	decodeBool       = flag.Bool("d", false, "Decode a string of numbers")
)

// isValidLanguage checks to see if the given string matches one of the language options
func isValidLanguage(s string) bool {
	switch s {
	case "en":
		return true
	case "el":
		return true
	default:
		return false
	}
}

// readFile opens a file matching the given filepath and returns
// the contents of the file as a string array for each line
func readFile(f string) []string {
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Failed to open file", f)
		fmt.Println(err)
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// rotateN handles when the user has the "-rot" flag enabled
func rotateN() {
	l := crypto.New(*lang)

	var args []string
	if *fileName != "" {
		args = readFile(*fileName)
	} else {
		args = flag.Args()
	}

	rotateString := strings.Join(args, " ")

	str, err := l.RotateN(rotateString, *rotateAmount)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

// polybiusSquare handles when the user has the "-ps" flag enabled
func polybiusSquare() {

	// Ensure that only encodeBool or decodeBool is set not both or neither
	if !*encodeBool && !*decodeBool || *encodeBool && *decodeBool {
		fmt.Println("Please select either \"-e\" or \"-d\", to encode or decode a string")
		os.Exit(-1)
	}

	var args []string
	if *fileName != "" {
		args = readFile(*fileName)
	} else {
		args = flag.Args()
	}

	l := crypto.New(*lang)
	if *encodeBool {
		for _, arg := range args {
			cypherText, err := l.PolybiusEncode(arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			fmt.Println(cypherText)
		}

	} else {
		for _, arg := range args {
			cypherText, err := l.PolybiusDecode(arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
			fmt.Println(cypherText)
		}
	}

}

func vigenereCypher() {

	// Ensure that only encodeBool or decodeBool is set not both or neither
	if !*encodeBool && !*decodeBool || *encodeBool && *decodeBool {
		fmt.Println("Please select either \"-e\" or \"-d\", to encode or decode a string")
		os.Exit(-1)
	}

	if *vigenereKey == "" {
		fmt.Println("Please use \"-k\" along with a key")
		os.Exit(-1)
	}

	var args []string
	if *fileName != "" {
		args = readFile(*fileName)
	} else {
		args = flag.Args()
	}

	singleArg := strings.Join(args, " ")

	l := crypto.New(*lang)

	if *encodeBool {
		cypherText, err := l.VigenereEncode(*vigenereKey, singleArg)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Println(cypherText)

	} else {
		plainText, err := l.VigenereDecode(*vigenereKey, singleArg)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Println(plainText)
	}
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Usage: ")
		fmt.Println("\t-l <language> -ps [-e | -d] [-f file_path | string]")
		fmt.Println("\t-l <language> -rot -n <int> [-f <file_path> | string]")
		fmt.Println("\t-l <language> -ver -k <string> [-f <file_path> | string]")
		fmt.Println("")
		flag.PrintDefaults()
	}

	if *lang == "" || !isValidLanguage(*lang) {
		fmt.Println("Please use the \"-l\" option and provide a valid language")
		os.Exit(-1)
	}

	if (*isRotateN && *isPolybiusSquare) ||
		(*isRotateN && *isVigenere) ||
		(*isPolybiusSquare && *isVigenere) {

		fmt.Println("Please only use one of the following:")
		fmt.Println("\t-ps")
		fmt.Println("\t-rot")
		fmt.Println("\t-ver")
		os.Exit(-1)
	}

	if *isRotateN {
		rotateN()
	} else if *isPolybiusSquare {
		polybiusSquare()
	} else if *isVigenere {
		vigenereCypher()
	}
}
