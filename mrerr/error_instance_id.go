package mrerr

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"
)

type (
	// ErrorIDGenerator - интерфейс генерации уникальных ID ошибок.
	ErrorIDGenerator interface {
		GenerateID() string
	}

	// IDGenerator - базовый генератор уникальных ID ошибок.
	IDGenerator struct{}
)

// NewIDGenerator - создаётся объект IDGenerator.
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// GenerateID - возвращает ID
// 'hex(unix time)' - 'hex(4 rand bytes)' -> 64e9c0f1-1e97228f
func (g *IDGenerator) GenerateID() string {
	var value [4]byte

	if _, err := rand.Read(value[:]); err != nil {
		value = [4]byte{0x0, 0xee, 0xee, 0x0} // suffix: 00eeee00
	}

	return strconv.FormatInt(time.Now().Unix(), 16) + "-" + hex.EncodeToString(value[:])
}
