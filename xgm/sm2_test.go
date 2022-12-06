package xgm

import (
	"testing"
)

func Test_SM2_GenerateKeyPair(t *testing.T) {
	// 生成密钥对
	publicKey, privateKey, err := SM2_GenerateKeyPair()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("publicKey", publicKey)
	t.Log("privateKey", privateKey)
	// 公钥加密
	cipher, err := SM2_PublicKey_Encrypt("11223344", publicKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("cipher", cipher)
	// 私钥解密
	plain, err := SM2_PrivateKey_Decrypt(cipher, privateKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("plain", plain)
}
