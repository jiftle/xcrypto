package xdes

import (
	"encoding/hex"
	"strings"
)

func TripleDes_S(key string, plain string, flag int) (out string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytPlain, err := hex.DecodeString(plain)
	if err != nil {
		return
	}
	bytOut, err := TripleDesEcb(bytKey, bytPlain, flag)
	if err != nil {
		return
	}
	out = hex.EncodeToString(bytOut)
	out = strings.ToUpper(out)
	return
}

func GetSubKey_S(key, dvs string) (out string, err error) {
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytDvs, err := hex.DecodeString(dvs)
	if err != nil {
		return
	}
	bytOut, err := GetSubKeyDES(bytKey, bytDvs)
	if err != nil {
		return
	}
	out = hex.EncodeToString(bytOut)
	out = strings.ToUpper(out)
	return
}
