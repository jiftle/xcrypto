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
	return
}
