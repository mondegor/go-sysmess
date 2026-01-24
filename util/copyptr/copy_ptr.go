package copyptr

import (
	"time"
)

// Time - возвращает копию значения времени или nil если значение равно nil или 0.
func Time(value *time.Time) *time.Time {
	if value == nil || value.IsZero() {
		return nil
	}

	c := *value

	return &c
}
