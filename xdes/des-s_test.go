package xdes

import (
	"fmt"
	"strings"
	"testing"
)

func TestTripleDes_S(t *testing.T) {
	key := "11111111111111111111111111111111"
	plain := "11111111111111111111111111111111"

	cipher, err := TripleDes_S(key, plain, CNT_MODE_ENCRYPT)
	if err != nil {
		t.Error(err)
	}

	if strings.EqualFold("F40379AB9E0EC533F40379AB9E0EC533", cipher) {
		t.Error("fail")
	}
	fmt.Println(cipher)
}
