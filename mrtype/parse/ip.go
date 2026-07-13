package parse

import (
	"net/netip"
	"strings"

	"github.com/mondegor/go-core/mrtype/errors"
)

const (
	typeIP   = "Ip"
	maxLenIP = 64
)

// IP - парсит строку в netip.Addr.
// Поддерживает форматы: IP-адрес (IPv4, IPv6) и адрес с портом ("1.2.3.4:80", "[2001:db8::1]:80").
// Адрес в формате IPv4-mapped IPv6 приводится к IPv4.
// Если значение пустое и required=true, возвращает ошибку.
// Если значение пустое и required=false, возвращает не заданный адрес.
func IP(value string, required bool) (netip.Addr, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return netip.Addr{}, errors.NewParamEmptyError(typeIP)
		}

		return netip.Addr{}, nil
	}

	if len(value) > maxLenIP {
		return netip.Addr{}, errors.NewParamLenMaxError(typeIP, maxLenIP)
	}

	// сначала проверяется формат чистого адреса, и только при неудаче - формат адреса с портом
	if addr, err := netip.ParseAddr(value); err == nil {
		return addr.Unmap(), nil
	}

	addrPort, err := netip.ParseAddrPort(value)
	if err != nil {
		return netip.Addr{}, errors.NewParamIncorrectError(typeIP, err)
	}

	return addrPort.Addr().Unmap(), nil
}
