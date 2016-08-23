package multibase

import (
	"encoding/hex"
	"fmt"

	b58 "github.com/jbenet/go-base58"
)

const (
	Base1        = '1'
	Base2        = '0'
	Base8        = '7'
	Base10       = '9'
	Base16       = 'f'
	Base58Flickr = 'Z'
	Base58BTC    = 'z'
)

var ErrUnsupportedEncoding = fmt.Errorf("selected encoding not supported")

func Encode(base int, data []byte) (string, error) {
	switch base {
	case Base58BTC:
		return string(Base58BTC) + b58.EncodeAlphabet(data, b58.BTCAlphabet), nil
	case Base16:
		return string(Base16) + hex.EncodeToString(data), nil
	default:
		return "", ErrUnsupportedEncoding
	}
}

func Decode(data string) (int, []byte, error) {
	switch data[0] {
	case Base58BTC:
		return Base58BTC, b58.DecodeAlphabet(data[1:], b58.BTCAlphabet), nil
	default:
		return -1, nil, ErrUnsupportedEncoding
	}
}
