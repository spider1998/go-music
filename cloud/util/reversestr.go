package util

//反转字符串
func ReverseStr(from string) string {
	runes := []rune(from)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseStringSlice(from []string) (result []string) {
	for i, _ := range from {
		result = append(result, from[len(from)-i-1])
	}
	return result
}

func ReverseIntSlice(from []int) (result []int) {
	for i, _ := range from {
		result = append(result, from[len(from)-i-1])
	}
	return result
}
