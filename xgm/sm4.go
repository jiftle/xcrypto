package xgm

import (
	"errors"
	"strconv"

	"github.com/tjfoc/gmsm/sm4"
)

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
