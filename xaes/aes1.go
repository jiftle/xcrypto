package xaes

import (
	"crypto/aes"
	"fmt"
)

const (
	// DECRYPT 解密模式
	DECRYPT = 0
	// ENCRYPT 加密模式
	ENCRYPT = 1
)

// Aes aes mode ecb
func Aes(data []byte, key []byte, flag int) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	dataLen := len(data)
	if dataLen%blockSize != 0 {
		err := fmt.Errorf("input length %d not multiple of %d bytes", dataLen, blockSize)
		return nil, err
	}
	src := data
	out := make([]byte, dataLen)
	dst := out
	for len(src) > 0 {
		switch flag {
		case ENCRYPT:
			block.Encrypt(dst, src[:blockSize])
		case DECRYPT:
			block.Decrypt(dst, src[:blockSize])
		}
		src = src[blockSize:]
		dst = dst[blockSize:]
	}
	return out, nil
}
