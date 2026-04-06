package casttype

import (
	"encoding/binary"
	"errors"
	"net"
)

// IP2uint - преобразует IPv4-адрес в числовое представление (uint32).
// Возвращает 0 без ошибки для пустого IP.
// Возвращает ошибку для IPv6 и некорректных адресов.
func IP2uint(ip net.IP) (uint32, error) {
	if len(ip) == 0 {
		return 0, nil
	}

	if ip4 := ip.To4(); ip4 != nil {
		return binary.BigEndian.Uint32(ip4), nil
	}

	if len(ip) == 16 {
		return 0, errors.New("no sane way to convert ipv6 into uint32")
	}

	return 0, errors.New("ip is incorrect")
}

// Uint2ip - преобразует числовое представление (uint32) в IPv4-адрес net.IP.
// Использует big-endian порядок байтов.
func Uint2ip(number uint32) net.IP {
	ip := make(net.IP, 4)

	binary.BigEndian.PutUint32(ip, number)

	return ip
}
