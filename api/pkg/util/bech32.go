package util

import (
	"encoding/hex"

	"github.com/btcsuite/btcutil/bech32"
)

func EncodeBech32(hrp string, hexStr string) (string, error) {
	// hex to bytes
	extStr := ""
	for i := 0; i < len(hexStr); i++ {
		extStr += "0" + string(hexStr[i])
	}
	bytes, err := hex.DecodeString(extStr)
	if err != nil {
		return "", err
	}
	// 4bit to 5bit
	conv, err := bech32.ConvertBits(bytes, 4, 5, true)
	if err != nil {
		return "", err
	}
	// encode
	encoded, err := bech32.Encode(hrp, conv)
	if err != nil {
		return "", err
	}
	return encoded, nil
}

func DecodeBech32(bech32Data string) (string, string, error) {
	// decode
	hrp, decoded, err := bech32.Decode(bech32Data)
	if err != nil {
		return "", "", err
	}
	// 5bit to 4bit
	conv, err := bech32.ConvertBits(decoded, 5, 4, false)
	if err != nil {
		return "", "", err
	}
	// bytes to hex
	hexStr := hex.EncodeToString(conv)
	compStr := ""
	for i := 1; i < len(hexStr); i += 2 {
		if len(compStr) == 64 {
			break
		}
		compStr += string(hexStr[i])
	}
	return hrp, compStr, nil
}

func EncodeNpub(hexStr string) (string, error) {
	return EncodeBech32("npub", hexStr)
}

func EncodeNsec(hexStr string) (string, error) {
	return EncodeBech32("nsec", hexStr)
}
