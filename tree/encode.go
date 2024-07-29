package tree

import "fmt"

type Lookup map[interface{}]string

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

func (t *HuffmanTree) Encode() string {
	res := ""

	for _, thing := range t.data {
		res += t.Lookup[thing]
	}

	return res
}

func (t *HuffmanTree) Decode(msg string) []interface{} {
	res := make([]interface{}, 0)
	curr := t.root

	for _, b := range msg {
		if b == '0' {
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
