package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type MD5Hash struct {
}

func NewMD5Hash() *MD5Hash {
	return &MD5Hash{}
}

func (h *MD5Hash) Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
