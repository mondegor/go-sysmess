package instance

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"strconv"
	"time"
)

// GenerateID - возвращает ID в формате 4R4K8OE9XZ-MWFN-4L6R используемый
// при создании конкретного экземпляра ошибки.
func GenerateID() string {
	var (
		value [28]byte // real 24 - 25
		rnd   [8]byte
	)

	buf := strconv.AppendInt(value[:0], time.Now().UnixNano(), 36)

	if _, err := rand.Read(rnd[:]); err != nil {
		rnd = [8]byte{0x0, 0x0, 0xee, 0xee, 0xee, 0xee, 0x0, 0x0}
	}

	buf = strconv.AppendUint(buf, binary.BigEndian.Uint64(rnd[:]), 36)

	buf[12] = '-'
	buf[17] = '-'

	return string(bytes.ToUpper(buf[2:22]))
}
