# Crypto Suite

![Build Status](https://github.com/GeistInDerSH/CryptoSuite/workflows/Build%20Status/badge.svg)

## Overview

The Crypto Suite repository contains a number of different classic cryptographic functions.

None of the cryptographic functions currently implemented are cryptographically secure and as such the original plaintext can be recovered. Do NOT use them for sensative data.


## Example Usage

### Rotate N

The `RotateN` function takes a plaintext string and an integer. The plaintext is then rotated the given number of characters. An example of the code can be seen the below.

Rotate N is not cryptographically secure.

```go
	c := crypto.New("en")

	cyphertext, err := c.RotateN("Cesar Cypher", 13)
	if err != nil {
		panic(err)
	}

	fmt.Println(cyphertext)
```

The result of this code is: `PRNFNE PLCURE`

### Polybius Square

A `Polybius Square` is a simple table for encoding plaintext strings into two integers based on their position in the table. For example: `Example` would become `15 53 11 32 35 31 15`. Further reading can be found [here](https://en.wikipedia.org/wiki/Polybius_square).

Polybius Square is not cryptographically secure.

#### Encode

The `PolybiusEncode` function takes a plaintext string and returns a string integers. Special characters, or whitespace characters will not be encoded, and as such will be lost.

```go
	c := crypto.New("en")

	cyphertext, err := c.PolybiusEncode("This is an example")
	if err != nil {
		panic(err)
	}

	fmt.Println(cyphertext)
```

The result of the above code is: `44 23 24 43 24 43 11 33 15 53 11 32 35 31 15`

#### Decode

The `PolybiusDecode` function takes an string of integers and returns a the plaintext string.

Note that the English Ploybius square uses the same integer for both "I" and "J"

```go
	c := crypto.New("en")

	cyphertext, err := c.PolybiusDecode("44 23 24 43 24 43 11 33 15 53 11 32 35 31 15")
	if err != nil {
		panic(err)
	}

	fmt.Println(cyphertext)
```

The result of the above code is: `THI/JSI/JSANEXAMPLE`

### Vigenere Cypher

The `Vigenere Cypher` is an algorithm the encode plaintext with a key. More information can be found [here](https://en.wikipedia.org/wiki/Vigenere_cipher).

Vigenere Cypher is not cryptographically secure.

#### Encode

The `VigenereEncode` function takes the key and plaintext and returns the cyphertext after performing the Vigenere Cypher algorithm.

```go
	c := crypto.New("en")

	cyphertext, err := c.VigenereEncode("Key", "Plain Text")
	if err != nil {
		panic(err)
	}

	fmt.Println(cyphertext)
```

The result of the above code is: `ZPYSR ROBR`

#### Decode

The `VigenereDecode` function takes the key and cyphertext and returns the plaintext after performing the Vigenere Cypher algorithm.

```go
	c := crypto.New("en")

	cyphertext, err := c.VigenereDecode("ZPYSR ROBR", "Key")
	if err != nil {
		panic(err)
	}

	fmt.Println(cyphertext)
```

The result of the above code is: `PLAIN TEXT`
