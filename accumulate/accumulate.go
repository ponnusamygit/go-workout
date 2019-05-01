package accumulate

// Accumulate function performs the given operator
func Accumulate(given []string, operator func(string) string) []string {
	result := []string{}
	for _, str := range given {
		result = append(result, operator(str))
	}
	return result
}
