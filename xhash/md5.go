package algorithm

import (
	"crypto/md5"

	"gitee.com/yctxkj/xcrypto/utility"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func MD5Endcode32(src string) string {
	MD5Inst := md5.New()
	b, _, _ := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder(), []byte(src))
	MD5Inst.Write(b)
	rst := MD5Inst.Sum(nil)
	dst := utility.Bin2HexString(rst)
	return dst
}
