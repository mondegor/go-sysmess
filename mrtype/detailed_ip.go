package mrtype

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

type (
	// DetailedIP - содержит информацию о настоящем IP и
	// об IP переданного прокси сервером через заголовки.
	DetailedIP struct {
		Real  net.IP `json:"real"`
		Proxy net.IP `json:"proxy"`
	}
)

// NewDetailedIP - создаёт DetailedIP из числовых представлений IPv4.
// Параметры:
//   - realIP - числовое представление реального IP клиента;
//   - proxyIP - числовое представление IP прокси-сервера;
//
// Если аргументы равны 0, соответствующие поля остаются nil.
func NewDetailedIP(realIP, proxyIP uint32) (ip DetailedIP) {
	if realIP > 0 {
		ip.Real = castUint2ip(realIP)
	}

	if proxyIP > 0 {
		ip.Proxy = castUint2ip(proxyIP)
	}

	return ip
}

// String - возвращает IP-адреса в виде строки.
// Формат: "real" или "real, proxy" (если proxy задан).
func (ip *DetailedIP) String() string {
	if len(ip.Proxy) == 0 || ip.Proxy.IsUnspecified() {
		return ip.Real.String()
	}

	return ip.Real.String() + ", " + ip.Proxy.String()
}

// ToUint - преобразует IP-адреса в числовое представление (uint32).
// Возвращает ошибку, если IP не является IPv4 или имеет неверную длину.
func (ip *DetailedIP) ToUint() (realIP, proxyIP uint32, err error) {
	realIP, err = castIP2uint(ip.Real)
	if err != nil {
		return 0, 0, fmt.Errorf("mrtype.ToUint: %w", err)
	}

	proxyIP, err = castIP2uint(ip.Proxy)
	if err != nil {
		return 0, 0, fmt.Errorf("mrtype.ToUint: %w", err)
	}

	return realIP, proxyIP, nil
}

// castIP2uint - преобразует IPv4-адрес в числовое представление (uint32).
// Возвращает 0 без ошибки для пустого IP.
// Возвращает ошибку для IPv6 и некорректных адресов.
func castIP2uint(ip net.IP) (uint32, error) {
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

// castUint2ip - преобразует числовое представление (uint32) в IPv4-адрес net.IP.
// Использует big-endian порядок байтов.
func castUint2ip(number uint32) net.IP {
	ip := make(net.IP, 4)

	binary.BigEndian.PutUint32(ip, number)

	return ip
}
