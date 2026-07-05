package mraccess

type (
	// rightsSet - предвычисленное множество прав (привилегий и разрешений)
	// группы ролей. Реализует RightsChecker с проверкой за O(1).
	rightsSet map[string]struct{}
)

// Has - сообщает, входит ли указанное право в множество.
func (s rightsSet) Has(name string) bool {
	_, ok := s[name]

	return ok
}
