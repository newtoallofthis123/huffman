package utils

func CalculateFreq(data []interface{}) map[interface{}]int {
	freq := make(map[interface{}]int)
	for _, d := range data {
		freq[d]++
	}
	return freq
}
