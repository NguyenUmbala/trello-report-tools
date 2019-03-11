package utils

import "time"

type UtilTime struct{}

func (*UtilTime) CompareTime(a *time.Time, b *time.Time) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil && a.String() == b.String() {
		return true
	}
	return false
}
