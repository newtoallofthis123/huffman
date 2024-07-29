package tree

type Lookup map[interface{}]string

func (t *HuffmanTree) makeLookUp() Lookup {
	lookup := make(Lookup)

	for n := t.root; !n.IsLeaf(); {
	}

	return lookup
}

func (t *HuffmanTree) Encode() []interface{} {
	return nil
}
