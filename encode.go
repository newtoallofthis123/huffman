package huffman

import "fmt"

type Lookup map[interface{}]string

// MakeLookUp creates the lookup table
// This is used internally, however can be used to make a lookup table of any data
func (t *HuffmanTree) MakeLookUp(n *HuffmanNode, path string, table *Lookup) {
	if n == nil {
		t.MakeLookUp(t.root, path, table)
	} else {
		if n.IsLeaf() {
			(*table)[n.data.(interface{})] = path
		} else {
			if n.left != nil {
				t.MakeLookUp(n.left, fmt.Sprintf("%s%c", path, '0'), table)
			}
			if n.right != nil {
				t.MakeLookUp(n.right, fmt.Sprintf("%s%c", path, '1'), table)
			}
		}
	}
}

// Encode encodes the data present in the huffman tree to a binary array
func (t *HuffmanTree) Encode() []uint8 {
	res := ""
	binary := make([]uint8, 0)

	for _, thing := range t.data {
		res += t.Lookup[thing]
	}

	for _, r := range res {
		if r == '0' {
			binary = append(binary, 0)
		} else {
			binary = append(binary, 1)
		}
	}

	return binary
}

// Decode decodes a binary array into the base data presented in a huffman tree
func (t *HuffmanTree) Decode(msg []uint8) []interface{} {
	res := make([]interface{}, 0)
	curr := t.root

	for _, b := range msg {
		if b == 0 {
			curr = curr.left
		} else {
			curr = curr.right
		}

		if curr.IsLeaf() {
			res = append(res, curr.data)
			curr = t.root
		}
	}

	return res
}
