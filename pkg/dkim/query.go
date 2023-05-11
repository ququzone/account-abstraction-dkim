package dkim

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"strings"
)

type queryResult struct {
	Verifier  Verifier
	KeyAlgo   string
	HashAlgos []string
	Notes     string
	Services  []string
	Flags     []string
}

type QueryMethod string

type txtLookupFunc func(domain string) ([]string, error)
type queryFunc func(domain, selector string, txtLookup txtLookupFunc) (*queryResult, error)

var queryMethods = map[QueryMethod]queryFunc{
	queryMethodDNSTXT: queryDNSTXT,
}

func queryDNSTXT(domain, selector string, txtLookup txtLookupFunc) (*queryResult, error) {
	var txts []string
	var err error
	if txtLookup != nil {
		txts, err = txtLookup(selector + "._domainkey." + domain)
	} else {
		txts, err = net.LookupTXT(selector + "._domainkey." + domain)
	}

	if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	txt := strings.Join(txts, "")

	return parsePublicKey(txt)
}

func parsePublicKey(s string) (*queryResult, error) {
	params, err := parseHeaderParams(s)
	if err != nil {
		return nil, err
	}

	res := new(queryResult)

	if v, ok := params["v"]; ok && v != "DKIM1" {
		return nil, errors.New("incompatible public key version")
	}

	p, ok := params["p"]
	if !ok {
		return nil, errors.New("key syntax error: missing public key data")
	}
	if p == "" {
		return nil, errors.New("key revoked")
	}
	p = strings.ReplaceAll(p, " ", "")
	b, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		return nil, err
	}
	switch params["k"] {
	case "rsa", "":
		pub, err := x509.ParsePKIXPublicKey(b)
		if err != nil {
			pub, err = x509.ParsePKCS1PublicKey(b)
			if err != nil {
				return nil, err
			}
		}
		rsaPub, ok := pub.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("key syntax error: not an RSA public key")
		}
		if rsaPub.Size()*8 < 1024 {
			return nil, errors.New(fmt.Sprintf("key is too short: want 1024 bits, has %v bits", rsaPub.Size()*8))
		}
		res.Verifier = &RsaVerifier{rsaPub}
		res.KeyAlgo = "rsa"
	default:
		return nil, errors.New("unsupported key algorithm")
	}

	if hashesStr, ok := params["h"]; ok {
		res.HashAlgos = parseTagList(hashesStr)
	}
	if notes, ok := params["n"]; ok {
		res.Notes = notes
	}
	if servicesStr, ok := params["s"]; ok {
		services := parseTagList(servicesStr)

		hasWildcard := false
		for _, s := range services {
			if s == "*" {
				hasWildcard = true
				break
			}
		}
		if !hasWildcard {
			res.Services = services
		}
	}
	if flagsStr, ok := params["t"]; ok {
		res.Flags = parseTagList(flagsStr)
	}

	return res, nil
}
