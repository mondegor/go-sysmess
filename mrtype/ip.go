package mrtype

import (
	"fmt"
	"net"

	"github.com/mondegor/go-sysmess/mrlib/casttype"
)

type (
	// DetailedIP - содержит информацию о настоящем IP и
	// об IP переданного прокси сервером через заголовки.
	DetailedIP struct {
		Real  net.IP
		Proxy net.IP
	}
)

// NewDetailedIP - comment method.
func NewDetailedIP(realIP, proxyIP uint32) (ip DetailedIP) {
	if realIP > 0 {
		ip.Real = casttype.Uint2ip(realIP)
	}

	if proxyIP > 0 {
		ip.Proxy = casttype.Uint2ip(proxyIP)
	}

	return ip
}

// String - возвращает IP в виде стоки.
func (ip *DetailedIP) String() string {
	if len(ip.Proxy) == 0 || ip.Proxy.IsUnspecified() {
		return ip.Real.String()
	}

	return ip.Real.String() + ", " + ip.Proxy.String()
}

// ToUint - возвращает IP в виде uint32 если это возможно.
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
