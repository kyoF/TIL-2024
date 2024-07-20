package valueobject

import (
	"crypto/sha256"
	"encoding/hex"
)

type Password struct {
	Value string
}

func NewPassword(password string) Password {
	return Password{
		Value: password,
	}
}

func (v *Password) Hash() string {
	hash := sha256.Sum256([]byte(v.Value))
	return hex.EncodeToString(hash[:])
}
