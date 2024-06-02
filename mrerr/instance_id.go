package mrerr

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"
)

// generateInstanceID - возвращает ID для текущей ошибки
// 'hex(unix time)' - 'hex(4 rand bytes)' -> 64e9c0f1-1e97228f
func generateInstanceID() string {
	var value [4]byte

	if _, err := rand.Read(value[:]); err != nil {
		value = [4]byte{0x0, 0xee, 0xee, 0x0} // suffix: 00eeee00
	}

	return strconv.FormatInt(time.Now().Unix(), 16) + "-" + hex.EncodeToString(value[:])
}
