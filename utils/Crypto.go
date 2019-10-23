package utils

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(str string) string {
	return fmt.Sprintf("%x\n", sha256.Sum256([]byte(str)))
}
