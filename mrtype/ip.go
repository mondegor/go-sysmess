package mrtype

import (
	"fmt"
	"net"

	"github.com/mondegor/go-sysmess/util/casttype"
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
		ip.Real = casttype.Uint2ip(realIP)
	}

	if proxyIP > 0 {
		ip.Proxy = casttype.Uint2ip(proxyIP)
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
	realIP, err = casttype.IP2uint(ip.Real)
	if err != nil {
		return 0, 0, fmt.Errorf("mrtype.ToUint: %w", err)
	}

	proxyIP, err = casttype.IP2uint(ip.Proxy)
	if err != nil {
		return 0, 0, fmt.Errorf("mrtype.ToUint: %w", err)
	}

	return realIP, proxyIP, nil
}
