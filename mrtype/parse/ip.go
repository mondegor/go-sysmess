package parse

import (
	"net"
	"strings"
)

const (
	typeIP   = "Ip"
	maxLenIP = 32
)

// IP - возвращает валидный IP адрес из указанной строки или ошибку, если парсинг не удался.
func IP(value string, required bool) (ip net.IP, err error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return nil, NewParamEmptyError(typeIP)
		}

		return nil, nil
	}

	if len(value) > maxLenIP {
		return nil, NewParamLenMaxError(typeIP, maxLenIP)
	}

	host := value

	if strings.Contains(value, ":") {
		host, _, err = net.SplitHostPort(value)
		if err != nil {
			return nil, NewParamIncorrectError(typeIP, err)
		}
	}

	if ip := net.ParseIP(host); ip != nil {
		return ip, nil
	}

	return nil, NewParamIncorrectError(typeIP, nil)
}
