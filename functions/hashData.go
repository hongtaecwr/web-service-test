package functions

import (
	"crypto/sha256"
	"fmt"
)

func HashData(strData string) string {
	h := sha256.New()
	h.Write([]byte(strData))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash
}
