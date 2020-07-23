package security

import (
	"crypto/sha1"
	"fmt"
)

// HashString hashes a string using SHA1 algorithm
func HashString(expression string) string {
	hash := sha1.New()
	hash.Write([]byte(expression))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
