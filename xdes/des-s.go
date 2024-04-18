package xdes

import (
	"encoding/hex"
	"fmt"
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

func PbocMac_S(key, iv, data string) (sout string, err error) {
	if len(key) != 32 {
		err = fmt.Errorf("key invalid")
		return
	}
	if len(iv) != 16 {
		err = fmt.Errorf("iv invalid")
		return
	}
	if len(data)%16 != 0 {
		err = fmt.Errorf("data invalid, must be a multiple of 8")
		return
	}
	bytKey, err := hex.DecodeString(key)
	if err != nil {
		return
	}
	bytIv, err := hex.DecodeString(iv)
	if err != nil {
		return
	}
	bytData, err := hex.DecodeString(data)
	if err != nil {
		return
	}

	bytMac, err := PbocMac(bytKey, bytIv, bytData, 4)
	if err != nil {
		return
	}
	sout = hex.EncodeToString(bytMac)
	sout = strings.ToUpper(sout)
	return
}
