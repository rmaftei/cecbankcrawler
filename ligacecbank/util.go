package ligacecbank



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
