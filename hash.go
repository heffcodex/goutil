package goutil

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5String(s string) string {
	sum := md5.Sum([]byte(s))

	return hex.EncodeToString(sum[:])
}
