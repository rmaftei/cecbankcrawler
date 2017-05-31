package ligacecbank

import (
	"strconv"
)

func isNotNumber(str string) bool {
	return !isNumber(str)
}

func isNumber(str string) bool {
	if _, err := strconv.Atoi(str); err == nil {
		return true
	}

	return false
}

func isNotBlank(str string) bool {
	return !isBlank(str)
}

func isBlank(str string) bool {
	if len(str) <= 0 {
		return true
	}

	for _, c := range str {
		if ' ' != c {
			return false
		}
	}

	return true
}
