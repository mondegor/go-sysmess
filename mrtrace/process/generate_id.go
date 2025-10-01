package process

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"strconv"
	"time"
)

// GenerateID - возвращает ID в формате C5R32KBPEIR-9AVDBA-SU2CN9-VDBA
// используемый при трейсинге.
func GenerateID() string {
	var (
		value [40]byte // real 34 - 38
		rnd   [8]byte
	)

	buf := strconv.AppendInt(value[:0], time.Now().UnixNano(), 36)

	if _, err := rand.Read(rnd[:]); err != nil {
		rnd = [8]byte{0x0, 0x0, 0xee, 0xee, 0xee, 0xee, 0x0, 0x0}
	}

	buf = strconv.AppendUint(buf, binary.BigEndian.Uint64(rnd[:]), 36)
	buf = strconv.AppendUint(buf, binary.BigEndian.Uint64(rnd[:]), 36)

	buf[12] = '-'
	buf[19] = '-'
	buf[26] = '-'

	return string(bytes.ToUpper(buf[1:31]))
}
