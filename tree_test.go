package huffman_test

import (
	"testing"

	"github.com/newtoallofthis123/huffman"
)

func TestHuffman(t *testing.T) {
	chars := []rune{'A', 'A', 'A', 'B', 'B', 'B', 'B', 'B', 'C', 'C', 'C', 'C', 'C', 'C', 'D', 'D', 'D', 'D', 'E', 'E'}
	interfaces := huffman.ConvertToInterface(chars)

	h := huffman.FromList(interfaces)
	res := h.Encode()
	r := h.Decode(res)

	for i, c := range r {
		if chars[i] != c {
			t.Error("Failed at spot", i)
		}
	}
}

func TestEncoding(t *testing.T) {
	chars := []rune{'A', 'A', 'A', 'B', 'B', 'B', 'B', 'B', 'C', 'C', 'C', 'C', 'C', 'C', 'D', 'D', 'D', 'D', 'E', 'E'}
	inters := huffman.ConvertToInterface(chars)

	h := huffman.FromList(inters)
	res := h.Encode()

	expected := "011011011101010101011111111111100000000010010"
	stringRes, _ := huffman.ConvertBinaryToString(res)

	if expected != stringRes {
		t.Error("Encode failed")
	}
}

func TestDecode(t *testing.T) {
	chars := []rune{'A', 'A', 'A', 'B', 'B', 'B', 'B', 'B', 'C', 'C', 'C', 'C', 'C', 'C', 'D', 'D', 'D', 'D', 'E', 'E'}
	inters := huffman.ConvertToInterface(chars)

	h := huffman.FromList(inters)

	decoded := "011011011101010101011111111111100000000010010"
	decodedBinary, _ := huffman.ConvertStringToBinary(decoded)

	toTest := h.Decode(decodedBinary)

	for i, c := range toTest {
		if chars[i] != c {
			t.Error("Decoded failed")
		}
	}
}
