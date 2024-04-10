package xgm

import (
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm3"
)

func SM3_Hash(data string) (hash string) {
	h := sm3.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)
	hash = hex.EncodeToString(sum)
	return
}
