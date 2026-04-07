package process

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

const (
	bufferLen   = 12 // math.Ceil(64 * math.Log(2) / math.Log(float64(baseNCharsLen)))
	idDash1Pos  = 12
	idDash2Pos  = 19
	idBufferLen = 24 // timestamp(12) + dash(1) + random1(6) + dash(1) + random2(4)
)

//nolint:gochecknoglobals
var (
	baseNChars    = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")        // символы для кодирования Base-N
	baseNCharsLen = uint64(len(baseNChars))                               // количество символов для кодирования
	rndIfError    = [...]byte{0x0, 0x0, 0xee, 0xee, 0xee, 0xee, 0x0, 0x0} // данные по умолчанию при ошибке
)

// GenerateID - генерирует уникальный идентификатор в формате XXXXXXXXXXXX-XXXXXX-XXXX.
// Используется для идентификации процессов при трассировке запросов.
func GenerateID() string {
	var (
		buf [bufferLen * 2]byte // (timestamp + random)
		rnd [len(rndIfError)]byte
	)

	pos := encodeBaseN(buf[:], uint64(time.Now().UnixNano())) //nolint:gosec

	if _, err := rand.Read(rnd[:]); err != nil {
		rnd = rndIfError
	}

	encodeBaseN(buf[pos:], binary.BigEndian.Uint64(rnd[:]))

	buf[idDash1Pos] = '-'
	buf[idDash2Pos] = '-'

	return string(buf[:idBufferLen]) // TTTTTTTTTTTT-XXXXXX-XXXX
}

func encodeBaseN(buf []byte, n uint64) (pos int) {
	var digits [bufferLen + 1]byte // 1 byte для защиты от переполнения

	// цифры переводятся в символы в обратном порядке
	for n > 0 {
		q := n / baseNCharsLen
		digits[pos] = baseNChars[n-q*baseNCharsLen]
		pos++
		n = q
	}

	if pos > bufferLen {
		pos = bufferLen
	}

	for i := 0; i < pos; i++ {
		buf[i] = digits[pos-1-i]
	}

	// остаток заполняется символом '0'
	for i := pos; i < bufferLen; i++ {
		buf[i] = baseNChars[0]
	}

	// pos всегда не больше bufferLen
	return pos
}
