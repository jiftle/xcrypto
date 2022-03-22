package padding

import "bytes"

func Padding_PKCS5(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func UnPadding_PKCS5(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 模式2（IEC 9797-1填充，或 PBOC2.0填充）：
//	以0x80和0x00构成的字节序列填充,0x80只出现在最前面,仅仅只出现且一次必须出现,剩下的字节以0x00填充.
func Padding_ISO7816_4(ciphertext []byte, blockSize int) []byte {
	var newPadText []byte
	padding := blockSize - len(ciphertext)%blockSize - 1
	padText := []byte{0x80}
	padText = append(padText, bytes.Repeat([]byte{byte(0)}, padding)...)
	newPadText = append(ciphertext, padText...)
	return newPadText
}

func UnPadding_ISO7816_4(plain []byte, blockSize int) []byte {
	var newPadText []byte
	lastIndex := bytes.LastIndex(plain, []byte{0x80})
	newPadText = plain[:lastIndex]
	return newPadText
}
