package utils

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func IsEmptyString(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func RandomString(length int) string {
	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVXYZabcdefghijklmnopqrstuvwxyz-_"
	result := ""

	for length > 0 {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}

		c := int(random.Int64())
		if c < len(charset) {
			result += string(charset[c])
			length--
		}
	}
	return result
}

func MappingGender(gender string) string {
	switch gender {
	case "F":
		return "female"
	case "M":
		return "male"
	case "O":
		return "other"
	default:
		return "unknown"
	}
}
