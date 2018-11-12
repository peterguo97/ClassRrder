package utils

import (
	"encoding/base64"
)

func CryptoPass(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
