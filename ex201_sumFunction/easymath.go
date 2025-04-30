package easymath

func sum(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	result := 0
	for _, value := range a {
		result += value
	}
	return result
}
