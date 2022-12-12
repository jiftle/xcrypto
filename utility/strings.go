package utility

import "strings"

// StringPadLeft 左补齐
func StringPadLeft(src string, paddingChar string, expc int) string {
	srcLen := len(src)
	if srcLen >= expc {
		return src[:expc]
	}
	return strings.Repeat(paddingChar, expc-srcLen) + src
}

// StringPadRight 右补齐
func StringPadRight(src string, paddingChar string, expc int) string {
	srcLen := len(src)
	if srcLen >= expc {
		return src[srcLen-expc : srcLen]
	}
	return src + strings.Repeat(paddingChar, expc-srcLen)
}

func StringTruncRigth(s string, l int, pad byte) string {
	ll := len(s)
	if ll > l {
		return s[ll-l:]
	}
	if ll < l {
		padding := StringPadLeft("", string(pad), l)
		return (padding + s)[ll:]
	}
	return s
}
