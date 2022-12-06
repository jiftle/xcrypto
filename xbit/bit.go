package xbit

import (
	"encoding/hex"
	"fmt"
)

func XOR(data1, data2 string) (out string, err error) {
	bytData1, err := hex.DecodeString(data1)
	if err != nil {
		return
	}
	bytData2, err := hex.DecodeString(data2)
	if err != nil {
		return
	}
	if len(bytData1) != len(bytData2) {
		err = fmt.Errorf("2块数据长度不相等")
		return
	}
	bytOut := make([]byte, len(bytData1))
	for i := 0; i < len(bytData1); i++ {
		bytOut[i] = bytData1[i] ^ bytData2[i]
	}
	out = hex.EncodeToString(bytOut)
	return
}
