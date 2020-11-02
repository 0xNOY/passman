package crypto

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func normalizeKeyForBCrypt(src []byte) []byte {
	h := sha3.Sum512(src)
	if len(src) <= 72 {
		return append(src, h[:]...)[:72]
	}
	return append(h[:], src...)[:72]
}

func HashKeyWithBCrypt(src []byte) ([]byte, error) {
	src = normalizeKeyForBCrypt(src)
	return bcrypt.GenerateFromPassword(src, bcrypt.DefaultCost)
}

func CompareHashAndKeyWithBCrypt(hashed, raw []byte) error {
	raw = normalizeKeyForBCrypt(raw)
	return bcrypt.CompareHashAndPassword(hashed, raw)
}
