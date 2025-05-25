package generate

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"strconv"
	"time"
)

// InstanceID - возвращает ID в формате DA54BADA-MO43-E7OA используемый
// при создании конкретного экземпляра ошибки.
func InstanceID() string {
	var (
		value [26]byte
		rnd   [8]byte
	)

	buf := strconv.AppendInt(value[:0], time.Now().UnixNano(), 36)

	if _, err := rand.Read(rnd[:]); err != nil {
		rnd = [8]byte{0x0, 0x0, 0xee, 0xee, 0xee, 0xee, 0x0, 0x0}
	}

	buf = strconv.AppendUint(buf, binary.BigEndian.Uint64(rnd[:]), 36)

	buf[8] = '-'
	buf[13] = '-'

	return string(bytes.ToUpper(buf[:18]))
}
