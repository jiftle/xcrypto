package xencoding

import "encoding/hex"

func HexStr2Utf8Str(data string) (out string, err error) {
	bytData, err := hex.DecodeString(data)
	if err != nil {
		return
	}
	out = string(bytData)
	return
}
