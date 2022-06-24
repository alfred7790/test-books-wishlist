package common

import (
	"crypto/md5"
	"encoding/hex"
)

func HashString(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
