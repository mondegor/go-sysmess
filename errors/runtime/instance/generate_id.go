package instance

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

const (
	bufferLen   = 13 // math.Ceil(64 * math.Log(2) / math.Log(float64(baseNCharsLen)))
	idOffset    = 1  // ведущие обрезанные символы
	idBufferLen = 20 // timestamp(10) + dash(1) + random1(4) + dash(1) + random2(4)
)

//nolint:gochecknoglobals
var (
	baseNChars    = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")        // символы для кодирования Base-N
	baseNCharsLen = uint64(len(baseNChars))                               // количество символов для кодирования
	rndIfError    = [...]byte{0x0, 0x0, 0xee, 0xee, 0xee, 0xee, 0x0, 0x0} // данные по умолчанию при ошибке
)

// GenerateID - генерирует уникальный идентификатор в формате XXXXXXXXXX-XXXX-XXXX.
// Используется для идентификации конкретных экземпляров runtime-ошибок.
func GenerateID() string {
	var (
		buf [bufferLen*2 - 1]byte // (timestamp + random)
		rnd [len(rndIfError)]byte
	)

	pos := encodeBaseN(buf[:], uint64(time.Now().UnixNano())) //nolint:gosec

	if _, err := rand.Read(rnd[:]); err != nil {
		rnd = rndIfError
	}

	// если timestamp занял 13 символов, сдвигаем начало random-секции на 1
	// влево, чтобы последний символ timestamp не вылез за дефис
	if pos > 11 {
		pos--
	}

	encodeBaseN(buf[pos:], binary.BigEndian.Uint64(rnd[:]))

	buf[11] = '-'
	buf[16] = '-'

	return string(buf[idOffset : idOffset+idBufferLen]) // xXXXXXXXXXX-XXXX-XXXXxxxx
}

func encodeBaseN(buf []byte, n uint64) (pos int) {
	var digits [bufferLen]byte

	// цифры переводятся в символы в обратном порядке
	for n > 0 {
		q := n / baseNCharsLen
		digits[pos] = baseNChars[n-q*baseNCharsLen]
		pos++
		n = q
	}

	for i := 0; i < pos; i++ {
		buf[i] = digits[pos-1-i]
	}

	return pos
}
