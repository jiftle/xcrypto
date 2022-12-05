package xgm

import (
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"strconv"

	"github.com/tjfoc/gmsm/sm4"
)

func PbocMac(data, key, iv string) (out string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytData, err := hex.DecodeString(data)
	if err != nil {
		return
	}
	byIv, err := hex.DecodeString(iv)
	if err != nil {
		return
	}
	bytOut, err := SM4PbocMac(bytKey, byIv, bytData, 4)
	if err != nil {
		return
	}
	out = hex.EncodeToString(bytOut)
	return
}

func DiversifyKey(key, dvs string) (newkey string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytDvs, err := hex.DecodeString(dvs)
	if err != nil {
		return
	}
	nk, err := SM4DiverseKey(bytDvs, 0, bytKey)
	if err != nil {
		return
	}
	newkey = hex.EncodeToString(nk)
	return
}

func Decrypt_CBC(data, key, iv string) (out string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytData, err := hex.DecodeString(data)
	if err != nil {
		return
	}
	byIv, err := hex.DecodeString(iv)
	if err != nil {
		return
	}
	bytOut, err := DecryptBytes_CBC(bytData, bytKey, byIv)
	if err != nil {
		return
	}
	out = hex.EncodeToString(bytOut)
	return
}

func DecryptBytes_CBC(data, key, iv []byte) (out []byte, err error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	out = make([]byte, len(data))
	blockMode.CryptBlocks(out, data)
	return
}

func Encrypt_CBC(data, key, iv string) (out string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytData, err := hex.DecodeString(data)
	if err != nil {
		return
	}
	byIv, err := hex.DecodeString(iv)
	if err != nil {
		return
	}
	bytOut, err := EncryptBytes_CBC(bytData, bytKey, byIv)
	if err != nil {
		return
	}
	out = hex.EncodeToString(bytOut)
	return
}

func EncryptBytes_CBC(data, key, iv []byte) (out []byte, err error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return
	}
	blockMode := cipher.NewCBCEncrypter(block, iv)
	if len(data)%blockMode.BlockSize() != 0 {
		err = errors.New("data lenght must be multi block")
		return
	}
	out = make([]byte, len(data))
	blockMode.CryptBlocks(out, data)
	return
}

func Encrypt_ECB(data, key string) (out string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytData, err := hex.DecodeString(data)
	if err != nil {
		return
	}
	bytOut, err := EncryptECB(bytData, bytKey)
	if err != nil {
		return
	}
	out = hex.EncodeToString(bytOut)
	return
}

func Decrypt_ECB(data, key string) (out string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytData, err := hex.DecodeString(data)
	if err != nil {
		return
	}
	bytOut, err := DecryptECB(bytData, bytKey)
	if err != nil {
		return
	}
	out = hex.EncodeToString(bytOut)
	return
}

func EncryptECB(data, key []byte) ([]byte, error) {
	var out []byte
	BlockSize := 16
	in := data
	if len(key) != BlockSize {
		return nil, errors.New("SM4: invalid key size " + strconv.Itoa(len(key)))
	}
	var inData []byte
	inData = in
	out = make([]byte, len(inData))
	c, err := sm4.NewCipher(key)
	if err != nil {
		return nil, errors.New("init new cipher fail")
	}
	for i := 0; i < len(inData)/16; i++ {
		in_tmp := inData[i*16 : i*16+16]
		out_tmp := make([]byte, 16)
		c.Encrypt(out_tmp, in_tmp)
		copy(out[i*16:i*16+16], out_tmp)
	}

	return out, nil
}

func DecryptECB(data, key []byte) ([]byte, error) {
	var out []byte
	BlockSize := 16
	in := data
	if len(key) != BlockSize {
		return nil, errors.New("SM4: invalid key size " + strconv.Itoa(len(key)))
	}
	var inData []byte
	inData = in
	out = make([]byte, len(inData))
	c, err := sm4.NewCipher(key)
	if err != nil {
		return nil, errors.New("init new cipher fail")
	}
	for i := 0; i < len(inData)/16; i++ {
		in_tmp := inData[i*16 : i*16+16]
		out_tmp := make([]byte, 16)
		c.Decrypt(out_tmp, in_tmp)
		copy(out[i*16:i*16+16], out_tmp)
	}

	return out, nil
}
