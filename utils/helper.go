package utils

func Max(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func Min(first, second int) int {
	if first > second {
		return second
	}
	return first
}
