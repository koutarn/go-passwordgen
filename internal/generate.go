package generate

import (
	"crypto/rand"
	"math/big"
)

func GeneratePassword(length int, charset string) (string, error) {
	var password []byte
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		charIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password = append(password, charset[charIndex.Int64()])
	}

	return string(password), nil
}