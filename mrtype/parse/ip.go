package parse

import (
	"net"
	"strings"

	"github.com/mondegor/go-sysmess/mrtype/errors"
)

const (
	typeIP   = "Ip"
	maxLenIP = 32
)

// IP - парсит строку в net.IP.
// Поддерживает форматы: чистый IP-адрес и "host:port" (извлекает host).
// Если значение пустое и required=true, возвращает ошибку.
// Если значение пустое и required=false, возвращает nil.
func IP(value string, required bool) (ip net.IP, err error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return nil, errors.NewParamEmptyError(typeIP)
		}

		return nil, nil
	}

	if len(value) > maxLenIP {
		return nil, errors.NewParamLenMaxError(typeIP, maxLenIP)
	}

	host := value

	if strings.Contains(value, ":") {
		host, _, err = net.SplitHostPort(value)
		if err != nil {
			return nil, errors.NewParamIncorrectError(typeIP, err)
		}
	}

	if ip := net.ParseIP(host); ip != nil {
		return ip, nil
	}

	return nil, errors.NewParamIncorrectError(typeIP, nil)
}
