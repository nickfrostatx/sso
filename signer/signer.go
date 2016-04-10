package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

const (
	salt = "sso-signer"
	sep  = "."
)

var (
	encoding        = base64.RawURLEncoding
	ErrBadSignature = errors.New("Bad signature")
	hashMethod      = sha256.New
)

func deriveKey(key []byte) []byte {
	h := hashMethod()
	h.Write([]byte(salt))
	h.Write(key)
	return h.Sum(nil)
}

type Signer struct {
	key []byte
}

func New(key []byte) *Signer {
	derived := deriveKey(key)
	return &Signer{
		key: derived,
	}
}

func (s *Signer) getSignature(data string) []byte {
	mac := hmac.New(hashMethod, s.key)
	mac.Write([]byte(data))
	return mac.Sum(nil)
}

func (s *Signer) Sign(data string) string {
	sig := s.getSignature(data)
	return data + sep + encoding.EncodeToString(sig)
}

func (s *Signer) Unsign(data string) (string, error) {
	parts := strings.SplitN(data, sep, 2)
	if len(parts) != 2 {
		return "", ErrBadSignature
	}
	sig, err := encoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}
	if !hmac.Equal(s.getSignature(parts[0]), sig) {
		return "", ErrBadSignature
	}
	return parts[0], nil
}
