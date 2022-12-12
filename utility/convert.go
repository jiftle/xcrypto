package utility

import (
	"encoding/hex"
	"strconv"
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	gbkDecoder = simplifiedchinese.GBK.NewDecoder()
	gbkEncoder = simplifiedchinese.GBK.NewEncoder()
)

// Bin2HexString 字节数组转16进制字符串
func Bin2HexString(binData []byte) string {
	if len(binData) == 0 {
		return ""
	}
	return hex.EncodeToString(binData)
}

// HexString2Bin 16进制字符串转字节数组
func HexString2Bin(strData string) []byte {
	b, _ := hex.DecodeString(strData)
	return b
}

// Bin2String 字节数组转asc字符串
func Bin2String(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

// String2Bin asc字符串转字节数组
func String2Bin(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bin2GBKString(b []byte) (string, error) {
	if len(b) == 0 {
		return "", nil
	}
	bo, err := gbkDecoder.Bytes(b)
	if err != nil {
		return "", err
	}
	return string(bo), nil
}

func GBKString2Bin(s string) ([]byte, error) {
	return gbkEncoder.Bytes([]byte(s))
}

// Bin2Uint32_BigEndian 字节数组转Uint32  大端
func Bin2Uint32_BigEndian(binData []byte) uint32 {
	b := make([]byte, 4)
	if len(binData) < 4 {
		copy(b[4-len(binData):4], binData)
	} else {
		copy(b, binData[:4])
	}
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

// Bin2Uint32_LittleEndian 字节数组转Uint32  小端
func Bin2Uint32_LittleEndian(binData []byte) uint32 {
	b := make([]byte, 4)
	if len(binData) < 4 {
		copy(b[4-len(binData):4], binData)
	} else {
		copy(b, binData[:4])
	}
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

// Bin2Uint16 字节数组转Uint16  大端
func Bin2Uint16(binData []byte) uint16 {
	b := make([]byte, 2)
	if len(binData) < 2 {
		copy(b[2-len(binData):2], binData)
	} else {
		copy(b, binData[:2])
	}
	return uint16(b[1]) | uint16(b[0])<<8
}

// Bin2Uint64 字节数组转Uint64  大端
func Bin2Uint64(binData []byte) uint64 {
	b := make([]byte, 8)
	if len(binData) < 8 {
		copy(b[8-len(binData):8], binData)
	} else {
		copy(b, binData[:8])
	}
	return uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
}

// Uint32ToBin_BigEndian Uint32转字节数组  大端
func Uint32ToBin_BigEndian(v uint32) []byte {
	b := make([]byte, 4)
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	return b
}

// Uint32ToBin_LittleEndian Uint32转字节数组  小端
func Uint32ToBin_LittleEndian(v uint32) []byte {
	b := make([]byte, 4)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
	return b
}

// Uint16ToBin Uint16转字节数组  大端
func Uint16ToBin(v uint16) []byte {
	b := make([]byte, 2)
	b[0] = byte(v >> 8)
	b[1] = byte(v)
	return b
}

// Uint64ToBin Uint64转字节数组  大端
func Uint64ToBin(v uint64) []byte {
	b := make([]byte, 8)
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
	return b
}

// IntToString Uint64转字符串  大端
func IntToString(n64 int64, radix int) string {
	return strconv.FormatInt(n64, radix)
}
