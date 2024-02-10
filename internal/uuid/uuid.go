package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

type UUID [16]byte

func New() (UUID, error) {
	var u UUID
	_, err := rand.Read(u[:])
	return u, err
}

func FromString(s string) (UUID, error) {
	var u UUID
	s = s[0:8] + s[9:13] + s[14:18] + s[19:23] + s[24:]
	hexBytes, err := hex.DecodeString(s)
	if err != nil {
		return u, err
	}

	if len(hexBytes) != 16 {
		return u, errors.New("invalid UUID length")
	}

	copy(u[:], hexBytes)
	return u, nil
}

func (u UUID) String() string {
	return hex.EncodeToString(u[:])
}
