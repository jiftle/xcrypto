package xgm

import (
	"crypto/rand"
	"errors"
	"math/big"

	"gitee.com/yctxkj/xcrypto/utility"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

var (
	uid                 = utility.String2Bin("xcrypto")
	errInvalidSignature = errors.New("invalid signature")
)

func GenerateKeyPair() ([]byte, error) {
	privKey, err := sm2.GenerateKey(nil) // 生成密钥对
	if err != nil {
		return nil, err
	}
	return utility.HexString2Bin(x509.WritePrivateKeyToHex(privKey)), nil
}

func SignWithPrivate(key []byte, msg []byte) ([]byte, error) {
	privKey, err := x509.ReadPrivateKeyFromHex(utility.Bin2HexString(key))
	if err != nil {
		return nil, err
	}
	r, s, err := sm2.Sm2Sign(privKey, msg, uid, rand.Reader)
	if err != nil {
		return nil, err
	}
	sig := make([]byte, 64)
	copy(sig, r.Bytes())
	copy(sig[33:], s.Bytes())
	return sig, nil
}

func VerifyWithPublic(key []byte, msg []byte, sig []byte) error {
	privKey, err := x509.ReadPrivateKeyFromHex(utility.Bin2HexString(key))
	if err != nil {
		return err
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])
	if !sm2.Sm2Verify(&privKey.PublicKey, msg, uid, r, s) {
		return errInvalidSignature
	}
	return nil
}
