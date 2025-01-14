package xgm

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/tjfoc/gmsm/sm3"
)

func Test_SM3_Hash(t *testing.T) {
	var data string
	var hash string

	data = "123456"
	h := sm3.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	hash = hex.EncodeToString(sum)
	fmt.Println(hash)
}

func Test_SM3_Hash_1(t *testing.T) {
	var data string
	var hash string

	data = "3E157E72266BC4B0ABAC624BB5492004DC81766A6D3E36E2DA3CCF9AB91C3280"
	h := sm3.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	hash = hex.EncodeToString(sum)
	fmt.Println(hash)
}
