package xpadding

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestPadding_PKCS5(t *testing.T) {
	plain, _ := hex.DecodeString("112233445566")

	byt := Padding_ISO7816_4(plain, 16)
	cipher := hex.EncodeToString(byt)
	fmt.Println(cipher)
}
