package dkim

import (
	"crypto"
	"crypto/rsa"
	"strings"
	"time"
)

type Verifier interface {
	Public() crypto.PublicKey
	Verify(hash crypto.Hash, hashed []byte, sig []byte) error
}

type Header struct {
	Version             string
	Algorithm           string
	Signature           []byte
	BodyHash            []byte
	Domain              string
	Headers             []string
	Auid                string
	QueryMethods        []string
	Selector            string
	SignatureTimestamp  time.Time
	SignatureExpiration time.Time
	HashAlgo            crypto.Hash
	Verifier            Verifier
	RawHeaderData       []byte
}

type RsaVerifier struct {
	*rsa.PublicKey
}

func (v *RsaVerifier) Public() crypto.PublicKey {
	return v.PublicKey
}

func (v *RsaVerifier) Verify(hash crypto.Hash, hashed, sig []byte) error {
	return rsa.VerifyPKCS1v15(v.PublicKey, hash, hashed, sig)
}

type headerPicker struct {
	h      header
	picked map[string]int
}

func newHeaderPicker(h header) *headerPicker {
	return &headerPicker{
		h:      h,
		picked: make(map[string]int),
	}
}

func (p *headerPicker) Pick(key string) string {
	at := p.picked[key]
	for i := len(p.h) - 1; i >= 0; i-- {
		kv := p.h[i]
		k, _ := parseHeaderField(kv)

		if !strings.EqualFold(k, key) {
			continue
		}

		if at == 0 {
			p.picked[key]++
			return kv
		}
		at--
	}

	return ""
}
