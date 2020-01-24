package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5 converts the user password to a hashed password using md5
func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(input))
	result := hex.EncodeToString(hash.Sum(nil))
	return result
}
