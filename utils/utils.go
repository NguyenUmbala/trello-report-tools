package utils

import (
	"strconv"
	"time"
)

type Utils struct{}

func (*Utils) CompareTime(a *time.Time, b *time.Time) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil && a.String() == b.String() {
		return true
	}
	return false
}

func (*Utils) GetTimeRealForDone(name string) int {
	l := len(name)
	time := ""
	for i := l - 1; i > 0; i-- {
		if string(name[i]) == "]" {
			i--
			for ; string(name[i]) != "["; i-- {
				time = string(name[i]) + string(time)
			}
			break
		}
		if i == 0 {
			time = "0"
		}
	}
	ret, _ := strconv.Atoi(time)
	return ret
}

func (*Utils) GetTimeGuessForDone(name string) int {
	l := len(name)
	time := ""
	for i := l - 1; i > 0; i-- {
		if string(name[i]) == ")" {
			i--
			for ; string(name[i]) != "("; i-- {
				time = string(name[i]) + string(time)
			}
			break
		}
		if i == 0 {
			time = "0"
		}
	}
	ret, _ := strconv.Atoi(time)
	return ret
}
