package copyptr

import (
	"time"
)

// Time - возвращает копию указателя на time.Time.
// Если value равен nil или нулевому времени, возвращает nil.
func Time(value *time.Time) *time.Time {
	if value == nil || value.IsZero() {
		return nil
	}

	c := *value

	return &c
}
