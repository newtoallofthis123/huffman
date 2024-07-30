package huffman_test

import (
	"reflect"
	"testing"

	"github.com/newtoallofthis123/huffman"
)

func TestBinaryStringToInts(t *testing.T) {
	tests := []struct {
		input    string
		expected []uint8
		hasError bool
	}{
		{"1010", []uint8{1, 0, 1, 0}, false},
		{"1111", []uint8{1, 1, 1, 1}, false},
		{"0000", []uint8{0, 0, 0, 0}, false},
		{"", []uint8{}, false},
		{"110201", nil, true},
		{"hello", nil, true},
	}

	for _, test := range tests {
		result, err := huffman.ConvertStringToBinary(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("Expected an error for input %q but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", test.input, err)
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For input %q, expected %v but got %v", test.input, test.expected, result)
			}
		}
	}
}

func TestConvertBinaryToString(t *testing.T) {
	tests := []struct {
		input    []uint8
		expected string
		hasError bool
	}{
		{[]uint8{1, 0, 1, 0}, "1010", false},
		{[]uint8{1, 1, 1, 1}, "1111", false},
		{[]uint8{0, 0, 0, 0}, "0000", false},
		{[]uint8{}, "", false},
		{[]uint8{1, 2, 0}, "", true},
		{[]uint8{3}, "", true},
	}

	for _, test := range tests {
		result, err := huffman.ConvertBinaryToString(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("Expected an error for input %v but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %v: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
			}
		}
	}
}
