package dkim

import (
	"crypto"
	"crypto/rsa"
	"math/big"
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
	fromIndex           int
	fromLength          int
	subjectIndex        int
	subjectLength       int
}

func (h *Header) Verify() error {
	hasher := h.HashAlgo.New()
	if _, err := hasher.Write(h.RawHeaderData); err != nil {
		return err
	}
	hashed := hasher.Sum(nil)
	return h.Verifier.Verify(h.HashAlgo, hashed, h.Signature)
}

func (h *Header) HeaderData() []byte {
	var meta []byte = make([]byte, 0)
	meta = append(meta, big.NewInt(int64(h.fromIndex)).Bytes()...)
	meta = append(meta, big.NewInt(int64(h.fromLength)).Bytes()...)
	meta = append(meta, big.NewInt(int64(h.subjectIndex)).Bytes()...)
	meta = append(meta, big.NewInt(int64(h.subjectLength)).Bytes()...)
	return append(meta, h.RawHeaderData...)
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
