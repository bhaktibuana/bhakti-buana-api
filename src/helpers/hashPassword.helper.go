package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword Helper
/*
 * @param password string
 * @returns string
 */
func HashPassword(password string) string {
	salt := "p@5s.13h@kt1.13u4n@"

	hasher := hmac.New(sha256.New, []byte(salt))
	hasher.Write([]byte(password))
	hashSum := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashSum)

	return hashString
}
