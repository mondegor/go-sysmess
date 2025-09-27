package casttype

import (
	"encoding/binary"
	"errors"
	"net"
)

// IP2uint - возвращает IP в виде числа или ошибку, если конвертация невозможна.
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

// Uint2ip - возвращает net.IP полученного из указанного целого числа.
func Uint2ip(number uint32) net.IP {
	ip := make(net.IP, 4)

	binary.BigEndian.PutUint32(ip, number)

	return ip
}
