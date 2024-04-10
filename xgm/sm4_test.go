package xgm

import (
	"encoding/hex"
	"strings"
	"testing"
)

func Test_SM4CryptEcb(t *testing.T) {
	// 密钥 C6E7DB6B460E3EB8A336964DC6746795
	// 明文 01020304050607080102030405060708
	// 密文 B6FF6C78FEF944883485E235301EF5C3
	key, _ := hex.DecodeString("C6E7DB6B460E3EB8A336964DC6746795")
	plain, _ := hex.DecodeString("01020304050607080102030405060708")
	cipherExpect := "B6FF6C78FEF944883485E235301EF5C3"
	cipher, _ := SM4CryptEcb(key, plain, 1)
	cipherVal := hex.EncodeToString(cipher)
	if !strings.EqualFold(cipherExpect, cipherVal) {
		t.Fatal("加密算法验证不通过")
	}
}

// func Test_SM4CryptEcb_Padding(t *testing.T) {
// 	// 密钥 C6E7DB6B460E3EB8A336964DC6746795
// 	// 明文 01020304050607080102030405060708
// 	// 密文 B6FF6C78FEF944883485E235301EF5C3
// 	key, _ := hex.DecodeString("C6E7DB6B460E3EB8A336964DC6746795")
// 	plain, _ := hex.DecodeString("0102030405060708")
// 	plain = Padding_ISO7816_4(plain, 16)
// 	cipherExpect := "B6FF6C78FEF944883485E235301EF5C3"
// 	cipher, _ := SM4CryptEcb(key, plain, 1)
// 	cipherVal := hex.EncodeToString(cipher)
// 	if !strings.EqualFold(cipherExpect, cipherVal) {
// 		t.Fatal("加密算法验证不通过")
// 	}
// }

// SM4CryptEcb SM4加解密CBC
// flag	0：解密，1：加密
func Test_SM4CryptCbc(t *testing.T) {
}

func Test_SM4PbocMac(t *testing.T) {
}

func Test_SM4DiverseKey(t *testing.T) {
	// 密钥 26ABD3D31EAD9C238B9414B0FE2465E6
	// 分散因子 030405060708090A
	// 结果 C6E7DB6B460E3EB8A336964DC6746795
	key, _ := hex.DecodeString("26ABD3D31EAD9C238B9414B0FE2465E6")
	dvs, _ := hex.DecodeString("030405060708090A")
	nk, _ := SM4DiverseKey(dvs, 0, key)
	sNk := hex.EncodeToString(nk)
	if !strings.EqualFold(sNk, "C6E7DB6B460E3EB8A336964DC6746795") {
		t.Fatal("分散验证不通过")
	}
}

func Test_SM4DiverseKey00(t *testing.T) {
	// 密钥 6AE7C80FD38E4D9ECAF6A557E52234FB
	// 分散因子 3344556677889900
	// 结果 0239BCD33FD0EACCB594502406405F71
	key, _ := hex.DecodeString("6AE7C80FD38E4D9ECAF6A557E52234FB")
	dvs, _ := hex.DecodeString("3344556677889900")
	nk, _ := SM4DiverseKey(dvs, 0, key)
	sNk := hex.EncodeToString(nk)
	if !strings.EqualFold(sNk, "0239BCD33FD0EACCB594502406405F71") {
		t.Fatal("分散验证不通过")
	}
}

// func Test_SM4_Encrypt_ECB(t *testing.T) {
// 	key := "C6E7DB6B460E3EB8A336964DC6746795"
// 	plain := "0102030405060708"
// 	cipher, err := SM4_Encrypt_ECB(key, plain)
// 	fmt.Println(cipher, err)
// }
