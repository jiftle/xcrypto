package gm

import (
	"encoding/hex"
	"testing"
)

func TestEncryptECB(t *testing.T) {
	sKey := "114D64906A39DCC3E908FD0A35EC12C7"
	sPlain := "4C232643F5C4238F719C9CF0F4968306"

	key, _ := hex.DecodeString(sKey)
	plain, _ := hex.DecodeString(sPlain)

	cipher, err := EncryptECB(plain, key)
	if err != nil {
		t.Logf("加密失败, %v", err)
		t.Fail()
	}

	t.Logf("密文: %v", hex.EncodeToString(cipher))
}

func TestDecryptECB(t *testing.T) {
	sKey := "114D64906A39DCC3E908FD0A35EC12C7"
	sPlain := "4F0A31B77AF5BC9F343814BA7C37B44C"

	key, _ := hex.DecodeString(sKey)
	plain, _ := hex.DecodeString(sPlain)

	t.Logf("密文: %v", sPlain)
	cipher, err := DecryptECB(plain, key)
	if err != nil {
		t.Logf("解密失败, %v", err)
		t.Fail()
	}

	t.Logf("明文: %v", hex.EncodeToString(cipher))
}
