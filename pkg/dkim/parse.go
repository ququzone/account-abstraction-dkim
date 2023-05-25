package dkim

import (
	"bufio"
	"crypto"
	"encoding/base64"
	"errors"
	"fmt"
	"net/textproto"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const (
	headerFieldName   = "DKIM-Signature"
	crlf              = "\r\n"
	queryMethodDNSTXT = "dns/txt"
)

type header []string

func readHeader(r *bufio.Reader) (header, error) {
	tr := textproto.NewReader(r)

	var h header
	for {
		l, err := tr.ReadLine()
		if err != nil {
			return h, fmt.Errorf("failed to read header: %v", err)
		}

		if len(l) == 0 {
			break
		} else if len(h) > 0 && (l[0] == ' ' || l[0] == '\t') {
			// This is a continuation line
			h[len(h)-1] += l + crlf
		} else {
			h = append(h, l+crlf)
		}
	}

	return h, nil
}

func parseHeaderField(s string) (k string, v string) {
	kv := strings.SplitN(s, ":", 2)
	k = strings.TrimSpace(kv[0])
	if len(kv) > 1 {
		v = strings.TrimSpace(kv[1])
	}
	return
}

func parseHeaderParams(s string) (map[string]string, error) {
	pairs := strings.Split(s, ";")
	params := make(map[string]string)
	for _, s := range pairs {
		kv := strings.SplitN(s, "=", 2)
		if len(kv) != 2 {
			if strings.TrimSpace(s) == "" {
				continue
			}
			return params, errors.New("dkim: malformed header params")
		}

		params[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}
	return params, nil
}

func parseTagList(s string) []string {
	tags := strings.Split(s, ":")
	for i, t := range tags {
		tags[i] = stripWhitespace(t)
	}
	return tags
}

func stripWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func parseTime(s string) (time.Time, error) {
	sec, err := strconv.ParseInt(stripWhitespace(s), 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0), nil
}

func decodeBase64String(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(stripWhitespace(s))
}

func removeSignature(s string) string {
	return regexp.MustCompile(`(b\s*=)[^;]+`).ReplaceAllString(s, "$1")
}

func parseCanonicalization(s string) (headerCan, bodyCan Canonicalization) {
	headerCan = CanonicalizationSimple
	bodyCan = CanonicalizationSimple

	cans := strings.SplitN(stripWhitespace(s), "/", 2)
	if cans[0] != "" {
		headerCan = Canonicalization(cans[0])
	}
	if len(cans) > 1 {
		bodyCan = Canonicalization(cans[1])
	}
	return
}

func Parse(data []byte, skipVerifier bool) (*Header, error) {
	header := &Header{}
	r := strings.NewReader(string(data))

	h, err := readHeader(bufio.NewReader(r))
	if err != nil {
		return nil, err
	}

	var signature string
	var index int
	for i, kv := range h {
		k, v := parseHeaderField(kv)
		if strings.EqualFold(k, headerFieldName) {
			signature = v
			index = i
			break
		}
	}

	params, err := parseHeaderParams(signature)
	if err != nil {
		return nil, err
	}

	header.Version = stripWhitespace(params["v"])
	header.Domain = stripWhitespace(params["d"])
	header.Selector = stripWhitespace(params["s"])
	header.Server = header.Selector + "." + header.Domain
	if i, ok := params["i"]; ok {
		header.Auid = stripWhitespace(i)
		if !strings.HasSuffix(header.Auid, "@"+header.Domain) && !strings.HasSuffix(header.Auid, "."+header.Domain) {
			return nil, errors.New("domain mismatch")
		}
	} else {
		header.Auid = "@" + header.Domain
	}

	headerKeys := parseTagList(params["h"])
	ok := false
	for _, k := range headerKeys {
		if strings.EqualFold(k, "from") {
			ok = true
			break
		}
	}
	if !ok {
		return nil, errors.New("from field not signed")
	}
	header.Headers = headerKeys

	if timeStr, ok := params["t"]; ok {
		t, err := parseTime(timeStr)
		if err != nil {
			return nil, err
		}
		header.SignatureTimestamp = t
	}
	if expiresStr, ok := params["x"]; ok {
		t, err := parseTime(expiresStr)
		if err != nil {
			return nil, err
		}
		header.SignatureExpiration = t
	}

	if !skipVerifier {
		methods := []string{string(queryMethodDNSTXT)}
		if methodsStr, ok := params["q"]; ok {
			methods = parseTagList(methodsStr)
		}
		var res *queryResult
		for _, method := range methods {
			if query, ok := queryMethods[QueryMethod(method)]; ok {
				res, err = query(header.Domain, stripWhitespace(params["s"]), nil)
				break
			}
		}
		if err != nil {
			return nil, err
		} else if res == nil {
			return nil, errors.New("unsupported public key query method")
		}
		header.Verifier = res.Verifier
	}

	header.Algorithm = stripWhitespace(params["a"])
	algos := strings.SplitN(header.Algorithm, "-", 2)
	if len(algos) != 2 {
		return nil, errors.New("malformed algorithm name")
	}
	keyAlgo := algos[0]
	hashAlgo := algos[1]

	if keyAlgo != "rsa" {
		return nil, errors.New("unsupported algorithm")
	}

	var hash crypto.Hash
	switch hashAlgo {
	case "sha256":
		hash = crypto.SHA256
	default:
		return nil, errors.New("unsupported hash algorithm")
	}
	header.HashAlgo = hash

	bodyHash, err := decodeBase64String(params["bh"])
	if err != nil {
		return nil, err
	}
	header.BodyHash = bodyHash

	sig, err := decodeBase64String(params["b"])
	if err != nil {
		return nil, err
	}
	header.Signature = sig

	headerData := make([]byte, 0)
	headerCan, _ := parseCanonicalization(params["c"])
	picker := newHeaderPicker(h)
	for _, key := range headerKeys {
		kv := picker.Pick(key)
		if kv == "" {
			continue
		}

		kv = canonicalizers[headerCan].CanonicalizeHeader(kv)
		if strings.EqualFold(key, "from") {
			header.fromIndex = len(headerData) + strings.Index(kv, "<") + 1 + 16
			header.fromLength = strings.Index(kv, ">") - strings.Index(kv, "<") - 1
		}
		if strings.EqualFold(key, "subject") {
			header.subjectIndex = len(headerData) + 8 + 16
			header.subjectLength = len(kv) - 10
			header.Subject = kv[8 : len(kv)-2]
		}
		headerData = append(headerData, []byte(kv)...)
	}
	canSigField := removeSignature(h[index])
	canSigField = canonicalizers[headerCan].CanonicalizeHeader(canSigField)
	canSigField = strings.TrimRight(canSigField, "\r\n")
	headerData = append(headerData, []byte(canSigField)...)
	header.RawHeaderData = headerData

	return header, nil
}
