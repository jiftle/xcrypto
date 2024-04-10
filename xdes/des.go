package xdes

import (
	"crypto/des"
	"errors"
	"strconv"
)

const (
	CNT_MODE_DECRYPT = 0
	CNT_MODE_ENCRYPT = 1
)

func DesEcb(keyValue []byte, txtValue []byte, flag int) ([]byte, error) {
	key := make([]byte, 8)
	if len(keyValue) >= 8 {
		copy(key, keyValue[:8])
	} else {
		err := errors.New("key length must be greater than 8 bytes")
		return nil, err
	}
	if len(txtValue)%8 != 0 {
		err := errors.New("input length not multiple of 8 bytes")
		return nil, err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	data := txtValue
	out := make([]byte, len(txtValue))
	dst := out
	for len(data) > 0 {
		if flag == CNT_MODE_ENCRYPT {
			block.Encrypt(dst, data[:blockSize])
		} else {
			block.Decrypt(dst, data[:blockSize])
		}
		data = data[blockSize:]
		dst = dst[blockSize:]
	}

	return out, nil
}

func TripleDesEcb(keyValue []byte, txtValue []byte, flag int) ([]byte, error) {
	key := make([]byte, 24)
	if len(keyValue) > 24 {
		copy(key, keyValue[:24])
	} else if len(keyValue) >= 16 {
		copy(key, keyValue[:16])
		copy(key[16:], keyValue[:8])
	} else {
		err := errors.New("key must be greater than 16")
		return nil, err
	}
	if len(txtValue)%8 != 0 {
		err := errors.New("input not multiple of 8 bytes")
		return nil, err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	data := txtValue
	out := make([]byte, len(txtValue))
	dst := out
	for len(data) > 0 {
		if flag == CNT_MODE_ENCRYPT {
			block.Encrypt(dst, data[:blockSize])
		} else {
			block.Decrypt(dst, data[:blockSize])
		}
		data = data[blockSize:]
		dst = dst[blockSize:]
	}

	return out, nil
}

// GetSubKeyDES 密钥分散
func GetSubKeyDES(ckKey []byte, dvsData []byte) ([]byte, error) {
	subData := make([]byte, 16)
	copy(subData[:8], dvsData[:8])
	for i := 0; i < 8; i++ {
		subData[8+i] = ^dvsData[i]
	}
	out, err := TripleDesEcb(ckKey, subData, CNT_MODE_ENCRYPT)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func PbocMac(key []byte, iv []byte, data []byte, macLen int) ([]byte, error) {
	if macLen > 8 || macLen < 0 {
		return nil, errors.New("invalid mac length " + strconv.Itoa(macLen))
	}
	blockcount := len(data) / 8
	macDataSize := blockcount * 8
	macData := make([]byte, macDataSize+8)
	copy(macData, data)
	if len(data)%8 != 0 {
		blockcount++
		macDataSize += 8
		macData[len(data)] = 0x80
	}

	blockData := make([]byte, 8)
	ivData := make([]byte, 8)
	if iv != nil {
		copy(ivData, iv[:8])
	}

	tmpMacData := macData
	var err error
	for i := 0; i < blockcount; i++ {
		for j := 0; j < 8; j++ {
			blockData[j] = tmpMacData[j] ^ ivData[j]
		}
		ivData, err = DesEcb(key[:8], blockData, CNT_MODE_ENCRYPT)
		if err != nil {
			return nil, err
		}
		tmpMacData = tmpMacData[8:]
	}
	blockData, err = DesEcb(key[8:], ivData, CNT_MODE_DECRYPT)
	if err != nil {
		return nil, err
	}
	ivData, err = DesEcb(key[:8], blockData, CNT_MODE_ENCRYPT)
	if err != nil {
		return nil, err
	}
	out := make([]byte, macLen)
	copy(out, ivData[:macLen])
	return out, nil
}
