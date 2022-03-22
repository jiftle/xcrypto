package xaes

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/encoding/gbase64"
)

var (
	content          = []byte("pibigstar")
	content_16, _    = gbase64.DecodeString("v1jqsGHId/H8onlVHR8Vaw==")
	content_24, _    = gbase64.DecodeString("0TXOaj5KMoLhNWmJ3lxY1A==")
	content_32, _    = gbase64.DecodeString("qM/Waw1kkWhrwzek24rCSA==")
	content_16_iv, _ = gbase64.DecodeString("DqQUXiHgW/XFb6Qs98+hrA==")
	content_32_iv, _ = gbase64.DecodeString("ZuLgAOii+lrD5KJoQ7yQ8Q==")
	// iv 长度必须等于blockSize，只能为16
	iv         = []byte("Hello My GoFrame")
	key_16     = []byte("1234567891234567")
	key_17     = []byte("12345678912345670")
	key_24     = []byte("123456789123456789123456")
	key_32     = []byte("12345678912345678912345678912345")
	keys       = []byte("12345678912345678912345678912346")
	key_err    = []byte("1234")
	key_32_err = []byte("1234567891234567891234567891234 ")
)

func Test_EncryptECB(t *testing.T) {
	key, _ := hex.DecodeString("11223344556677881122334455667788")
	plain, _ := hex.DecodeString("11223344556677881122334455667788")
	cipher, err := Encrypt_ECB(plain, key)
	if err != nil {
		t.Fail()
	}
	sCipher := hex.EncodeToString(cipher)
	fmt.Println(sCipher)

	//func Encrypt_ECB(plain, key []byte) ([]byte, error) {
	//t.Assert(cipher, []byte(content_16))
}

func Test_DecryptECB(t *testing.T) {
	key, _ := hex.DecodeString("11223344556677881122334455667788")
	plain, _ := hex.DecodeString("00f59e5c63934fd0efd90b057d1a2ad1")
	cipher, err := Decrypt_ECB(plain, key)
	if err != nil {
		t.Fail()
	}
	sCipher := hex.EncodeToString(cipher)
	fmt.Println(sCipher)

	//func Encrypt_ECB(plain, key []byte) ([]byte, error) {
	//t.Assert(cipher, []byte(content_16))
}

func TestEncryptAndDecrypt(t *testing.T) {
	//
	//0010112233445566778899AABBCCDDEEFF00
	sOrgPlain := "0010112233445566778899AABBCCDDEEFF00"
	sKey := "901E528CF1BDF6847A366F282F4C6A1C"
	//sOrgPlain := "11223344556677881122334455667788"
	//sKey := "11223344556677881122334455667788"

	key, _ := hex.DecodeString(sKey)
	plain, _ := hex.DecodeString(sOrgPlain)

	t.Logf("plain: %v", sOrgPlain)
	cipher, err := EncryptECB_Pad(plain, key)
	if err != nil {
		t.Logf("加密失败, %v", err)
		t.Fail()
	}
	sCipher := hex.EncodeToString(cipher)
	t.Logf("cipher: %v", sCipher)

	plain, _ = DecryptECB_Pad(cipher, key)
	sPlain := hex.EncodeToString(plain)
	t.Logf("plain: %v", sPlain)
	if sPlain != sOrgPlain {
		t.Fail()
	}

}
