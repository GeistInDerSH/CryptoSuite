package language

import (
	"github.com/GeistInDerSH/CryptoSuite/language/el"
	"github.com/GeistInDerSH/CryptoSuite/language/en"
)

type Crypto interface {
	Init()
	PolybiusEncode(string) (string, error)
	PolybiusDecode(string) (string, error)
	RotateN(string, int) (string, error)
	VigenereEncode(key, plaintext string) (string, error)
	VigenereDecode(key, cyphertext string) (string, error)
}

func New(s string) Crypto {
	var l Crypto
	switch s {
	case "en":
		l = &en.EN{}
	case "el":
		l = &el.EL{}
	default:
		l = nil
	}

	l.Init()

	return l
}
