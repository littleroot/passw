// Package passw provides a password generator using crypto/rand.
package passw

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// the following characters are skipped since they are sometimes difficult to
// distinguish from each other: 0, 1, I, O, i, l, o
var chars = [...]byte{
	'2', '3', '4',
	'5', '6', '7', '8', '9',

	'A', 'B', 'C', 'D', 'E',
	'F', 'G', 'H', 'J',
	'K', 'L', 'M', 'N',
	'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y',
	'Z',

	'a', 'b', 'c', 'd', 'e',
	'f', 'g', 'h', 'j',
	'k', 'm', 'n',
	'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y',
	'z',
}

var charsLen = big.NewInt(int64(len(chars)))

const (
	numParts      = 4
	partLen       = 3
	partSeparator = '-'
)

// Generate generates a new password.
func Generate() (string, error) {
	var buf strings.Builder

	for i := 0; i < numParts; i++ {
		part, err := generatePart()
		if err != nil {
			return "", err
		}
		buf.Write(part)

		if i != numParts-1 {
			buf.WriteByte(partSeparator)
		}
	}

	return buf.String(), nil
}

func generatePart() ([]byte, error) {
	p := make([]byte, partLen)

	for j := 0; j < partLen; j++ {
		var err error
		p[j], err = randomChar()
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func randomChar() (byte, error) {
	idx, err := rand.Int(rand.Reader, charsLen)
	if err != nil {
		return 0, err
	}
	return chars[idx.Int64()], nil
}
