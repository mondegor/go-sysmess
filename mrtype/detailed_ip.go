package mrtype

import (
	"net/netip"
)

type (
	// DetailedIP - содержит информацию о настоящем IP и
	// об IP переданного прокси сервером через заголовки.
	// Поддерживаются адреса IPv4 и IPv6 в едином представлении netip.Addr,
	// не заданный адрес - это нулевое (невалидное) значение netip.Addr.
	DetailedIP struct {
		Real  netip.Addr `json:"real"`
		Proxy netip.Addr `json:"proxy"`
	}
)

// NewIP - создаёт DetailedIP с указанным реальным IP клиента (IPv4 или IPv6).
// Параметры:
//   - realIP - реальный IP клиента;
//
// Если адрес не задан, поле Real остаётся не заданным.
func NewIP(realIP netip.Addr) DetailedIP {
	return DetailedIP{
		Real: realIP.Unmap(),
	}
}

// NewDetailedIP - создаёт DetailedIP с указанными IP клиента и прокси-сервера (IPv4 или IPv6).
// Параметры:
//   - realIP - реальный IP клиента;
//   - proxyIP - IP прокси-сервера;
//
// Если адрес не задан, соответствующее поле остаётся не заданным.
func NewDetailedIP(realIP, proxyIP netip.Addr) DetailedIP {
	return DetailedIP{
		Real:  realIP.Unmap(),
		Proxy: proxyIP.Unmap(),
	}
}

// String - возвращает IP-адреса в виде строки.
// Формат: "real" или "real, proxy" (если proxy задан). Не заданный IP даёт пустую строку.
// Если задан только proxy, real подставляется как "0" ("0, proxy").
func (ip DetailedIP) String() string {
	realStr := ipToString(ip.Real)

	if !ip.Proxy.IsValid() || ip.Proxy.IsUnspecified() {
		return realStr
	}

	if realStr == "" {
		realStr = "0"
	}

	return realStr + ", " + ip.Proxy.String()
}

// ipToString - возвращает строковое представление IP или "" для не заданного адреса.
func ipToString(ip netip.Addr) string {
	if !ip.IsValid() {
		return ""
	}

	return ip.String()
}
