package huffman

import "fmt"

// CalculateFreq calculates the frequency of each element in the given data.
// Can be useful for recalibrating the Huffman tree.
func CalculateFreq(data []interface{}) map[interface{}]int {
	freq := make(map[interface{}]int)
	for _, d := range data {
		freq[d]++
	}
	return freq
}

// ConvertToInterface converts a slice of type T to a slice of interface{}.
// Since most of the functions in the package work with interface{}, this function is useful for converting data.
// It should be used with caution if the data is not of type T.
func ConvertToInterface[T any](slice []T) []interface{} {
	interfaces := make([]interface{}, len(slice))
	for i, v := range slice {
		interfaces[i] = v
	}
	return interfaces
}

// ConvertStringToBinary converts a string representation of binary data into a byte slice.
func ConvertStringToBinary(s string) ([]uint8, error) {
	for i := 0; i < len(s); i++ {
		if s[i] != '0' && s[i] != '1' {
			return nil, fmt.Errorf("invalid binary string: %s", s)
		}
	}

	ints := make([]uint8, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ints[i] = 1
		} else {
			ints[i] = 0
		}
	}

	return ints, nil
}

// ConvertBinaryToString converts a binary representation of a byte slice into a string.
// Useful for easy printing of binary data.
func ConvertBinaryToString(b []uint8) (string, error) {
	res := ""

	for _, c := range b {
		if c == 0 {
			res += "0"
		} else if c == 1 {
			res += "1"
		} else {
			return "", fmt.Errorf("invalid binary")
		}
	}

	return res, nil
}
