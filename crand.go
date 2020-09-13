package crand

import (
	"crypto/rand"
	"golang.org/x/xerrors"
	"math/big"
)

const (
	defaultCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

// String generates a crypto secure random string of length size using the default charset.
func String(size int) (string, error) {
	return generate(size, defaultCharset)
}

// StringWithCharset generates a crypto secure random string of length size using the specified charset.
func StringWithCharset(size int, charset string) (string, error) {
	return generate(size, charset)
}

func generate(size int, charset string) (string, error) {
	b := make([]byte, size)

	for i := 0; i < size; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", xerrors.Errorf("generate string: %w", err)
		}

		b[i] = charset[n.Int64()]
	}

	return string(b), nil
}
