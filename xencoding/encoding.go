package xencoding

import (
	"encoding/hex"
	"strings"
)

func HexStr2Utf8Str(data string) (out string, err error) {
	bytData, err := hex.DecodeString(data)
	if err != nil {
		return
	}
	out = string(bytData)
	return
}

func Utf8Str2HexStr(data string) (out string) {
	bytData := []byte(data)
	out = hex.EncodeToString(bytData)
	out = strings.ToUpper(out)
	return
}
