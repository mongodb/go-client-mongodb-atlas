package generators

import (
	"crypto/rand"
	"math/big"
)

func RandomInteger(min, max int) (int, error) {
	offset, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, err
	}
	return min + int(offset.Int64()), err
}

func RandomASCIIString(length int) (string, error) {
	const asciiMax = 127
	result := ""
	for {
		if len(result) >= length {
			return result, nil
		}
		num, err := rand.Int(rand.Reader, big.NewInt(int64(asciiMax)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		// Make sure that the number/byte/letter is inside
		// the range of printable ASCII characters (excluding space and DEL)
		if n > 64 && n < asciiMax {
			result += string(rune(n))
		}
	}
}

func UniqueName(prefix string) (string, error) {
	const asciiMax = 122
	const length = 5
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(asciiMax)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		// Make sure that the number/byte/letter is inside
		// the range of printable ASCII characters (excluding space and DEL)
		if n > 97 && n < asciiMax {
			prefix += string(rune(n))
		}
	}
	return prefix, nil
}
