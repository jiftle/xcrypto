package xrand

import (
	"bytes"
	"math/rand"
	"time"

	"gitee.com/yctxkj/xcrypto/utility"
)

// GetRandString 获取随机数字符串（字符长度）
func GetRandStr(length int) string {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var buffer = bytes.Buffer{}
	var value int32 = 0
	for buffer.Len() < length*2 {
		value = rand.Int31()
		buffer.WriteString(utility.IntToString(int64(value), 16))
	}
	str := buffer.String()
	return str[length/2 : length/2+length]
}
