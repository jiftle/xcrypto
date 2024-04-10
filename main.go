package main

import (
	"encoding/hex"
	"fmt"

	"github.com/jiftle/xcrypto/xaes"
)

func main() {
	fmt.Println("vim-go")

	//plain, _ := hex.DecodeString("112233445566")

	//byt := padding.Padding_ISO7816_4(plain, 16)
	//cipher := hex.EncodeToString(byt)
	//fmt.Println(cipher)

	sOrgPlain := "0010112233445566778899AABBCCDDEEFF00"
	sKey := "901E528CF1BDF6847A366F282F4C6A1C"

	key, _ := hex.DecodeString(sKey)
	plain, _ := hex.DecodeString(sOrgPlain)

	cipher, err := xaes.EncryptECB_Pad(plain, key)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cipher)

}
