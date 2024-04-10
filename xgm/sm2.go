package xgm

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func SM2_GenerateKeyPair() (publicKey string, privateKey string, err error) {
	privKey, err := sm2.GenerateKey(nil)
	if err != nil {
		return
	}
	privateKey = x509.WritePrivateKeyToHex(privKey)
	publicKey = x509.WritePublicKeyToHex(&privKey.PublicKey)
	return
}

func SM2_PublicKey_Encrypt(plain string, publicKey string) (cipher string, err error) {
	pub, err := x509.ReadPublicKeyFromHex(publicKey)
	if err != nil {
		return
	}
	bytPlan, err := hex.DecodeString(plain)
	if err != nil {
		return
	}
	d, err := sm2.Encrypt(pub, bytPlan, rand.Reader, sm2.C1C3C2)
	if err != nil {
		return
	}
	cipher = hex.EncodeToString(d)
	return
}

func SM2_PrivateKey_Decrypt(cipher string, privKey string) (plain string, err error) {
	priv, err := x509.ReadPrivateKeyFromHex(privKey)
	if err != nil {
		return
	}
	bytCipher, err := hex.DecodeString(cipher)
	if err != nil {
		return
	}
	d, err := sm2.Decrypt(priv, bytCipher, sm2.C1C3C2)
	if err != nil {
		return
	}
	plain = hex.EncodeToString(d)
	return
}
