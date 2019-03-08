package services

import "strconv"

// Handle string
// Get time in name of card
// Example: "Demo(2)[3]" => get (2) and [3]

func GetRealTimeOfDone(name string) int {
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

func GetTimeGuessForDone(name string) int {
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
