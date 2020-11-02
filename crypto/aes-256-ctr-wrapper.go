package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/sha3"
)

func normalizeKeyForAES(src []byte) []byte {
	h := sha3.Sum256(src)
	if len(src) <= 32 {
		return append(src, h[:]...)[:32]
	}
	return h[:]
}

func Encrypt(key, src []byte) ([]byte, error) {
	key = normalizeKeyForAES(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	encrypted := make([]byte, aes.BlockSize+len(src))
	iv := encrypted[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], src)

	return encrypted, nil
}

func Decrypt(key, src []byte) ([]byte, error) {
	key = normalizeKeyForAES(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(src)-aes.BlockSize)
	stream := cipher.NewCTR(block, src[:aes.BlockSize])
	stream.XORKeyStream(decrypted, src[aes.BlockSize:])

	return decrypted, nil
}
