package xaes

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"

	"gitee.com/yctxkj/xcrypto/xpadding"
)

func EncryptECB_Pad(plain, key []byte) ([]byte, error) {
	bytPlain := xpadding.Padding_ISO7816_4(plain, 16)
	bytCipher, err := Encrypt_ECB(bytPlain, key)
	if err != nil {
		return nil, err
	}
	return bytCipher, err
}

func DecryptECB_Pad(plain, key []byte) ([]byte, error) {
	bytCipher, err := Decrypt_ECB(plain, key)
	if err != nil {
		return nil, err
	}
	bytCipherNoPad := xpadding.UnPadding_ISO7816_4(bytCipher, 16)
	return bytCipherNoPad, err
}

func Encrypt_ECB(plain, key []byte) ([]byte, error) {
	bytCipher := make([]byte, 0)

	//fmt.Printf("plain len= %v\n", len(plain))
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := cipher.BlockSize()
	if len(plain)%blockSize != 0 {
		return nil, errors.New("data isn't block size count!")
	}

	//fmt.Printf("blockSize: %v\n", blockSize)
	cipherLen := len(plain)
	bytCipher = make([]byte, cipherLen)
	for i := 0; i < cipherLen; i += blockSize {
		//fmt.Printf("bytCipher[%d:%d], bytPlain[%d:%d]\n", i, i+blockSize, i, i+blockSize)
		cipher.Encrypt(bytCipher[i:i+blockSize], plain[i:i+blockSize])
	}

	return bytCipher, nil
}

func Decrypt_ECB(bytCipher, key []byte) ([]byte, error) {
	bytPlain := make([]byte, 0)

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := cipher.BlockSize()
	if len(bytCipher)%blockSize != 0 {
		return nil, errors.New("data isn't block size interger count!")
	}

	cipherLen := len(bytCipher)
	bytPlain = make([]byte, cipherLen)
	for i := 0; i < cipherLen; i += blockSize {
		cipher.Decrypt(bytPlain[i:i+blockSize], bytCipher[i:i+blockSize])
	}

	return bytPlain, nil
}

// ================== CBC ========================

func AES_Encrypt_CBC(plantText, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plantText = xpadding.Padding_ISO7816_4(plantText, block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func AES_Decrypt_CBC(ciphertext, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	length := len(ciphertext)
	if length%block.BlockSize() != 0 {
		return nil, errors.New("data invalid")
	}

	blockModel := cipher.NewCBCDecrypter(block, iv)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = xpadding.UnPadding_ISO7816_4(plantText, block.BlockSize())
	return plantText, nil
}
